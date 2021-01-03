package ginlogger

import (
	"bytes"
	"fmt"
	"github.com/alonelegion/account_storage_mongo/pkg/ginbodywriter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

func Logger(logger logrus.FieldLogger) gin.HandlerFunc {
	hostname := getHostname()

	return func(c *gin.Context) {
		blw := ginbodywriter.NewWriter(c.Writer)
		c.Writer = blw

		path, method, referrer, userAgent := getContextData(c)
		if strings.Contains(path, "__debug") {
			return
		}

		var requestBody string
		c.Request.Body, requestBody = copyBody(c.Request.Body)

		latency := calculateLatency(c.Next)

		statusCode := c.Writer.Status()

		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":       hostname,
			"statusCode":     statusCode,
			"latency":        latency, // time to process
			"clientIP":       c.ClientIP(),
			"method":         method,
			"path":           path,
			"referer":        referrer,
			"dataLength":     dataLength,
			"userAgent":      userAgent,
			"requestBody":    requestBody,
			"requestHeaders": c.Request.Header,
			"responseBody":   blw.Body.String(),
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s] \"%s %s\" %d",
				time.Now().Format(timeFormat),
				c.Request.Method,
				path,
				statusCode,
			)
			switch {
			case statusCode >= http.StatusInternalServerError:
				entry.Error(msg)
			case statusCode >= http.StatusMultipleChoices:
				entry.Warn(msg)
			default:
				entry.Info(msg)
			}
		}
	}
}

func calculateLatency(f func()) int64 {
	start := time.Now()

	f()

	return time.Since(start).Milliseconds()
}

func copyBody(reader io.Reader) (io.ReadCloser, string) {
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return ioutil.NopCloser(bytes.NewBufferString("")), ""
	}

	return ioutil.NopCloser(bytes.NewBuffer(bodyBytes)), string(bodyBytes)
}

func getContextData(c *gin.Context) (string, string, string, string) {
	return c.Request.URL.Path,
		c.Request.Method,
		c.Request.Referer(),
		c.Request.UserAgent()
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return hostname
}
