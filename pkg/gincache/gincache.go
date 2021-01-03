package gincache

import (
	"encoding/json"
	"github.com/alonelegion/account_storage_mongo/internal/application"
	"github.com/alonelegion/account_storage_mongo/pkg/ginbodywriter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CacheResponse(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		gbw := ginbodywriter.NewWriter(c.Writer)
		c.Writer = gbw

		cacheBytes, err := application.Storage().Get("route_cache_" + c.Request.RequestURI)
		if err == nil {
			var cache Cache

			if err := json.Unmarshal(cacheBytes, &cache); err != nil {
				logger.WithError(err).Error("error when unmarshaling cached data, skip cache.")

				goto NoCacheBehavior
			}

			c.Header("XCached", "true")
			c.Data(cache.Status, "application/json", cache.Body)
			c.Abort()

			return
		}

	NoCacheBehavior:
		c.Next()

		cacheSettings, set := c.Get("cache")
		if !set {
			return
		}

		setting, ok := cacheSettings.(Setting)
		if !ok {
			logger.Error("wrong settings given")

			return
		}

		cacheBytes, err = Cache{
			Status: c.Writer.Status(),
			Body:   gbw.Bytes(),
		}.ToBytes()

		if err != nil {
			logger.WithError(err).Error("error when marshaling data for caching, skip cache.")

			return
		}

		application.Storage().Set("route_cache_"+c.Request.RequestURI, cacheBytes, setting.CacheDuration)
	}
}
