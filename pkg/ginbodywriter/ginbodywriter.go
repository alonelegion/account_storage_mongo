package ginbodywriter

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type GinBodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func NewWriter(responseWriter gin.ResponseWriter) *GinBodyWriter {
	return &GinBodyWriter{
		Body:           bytes.NewBufferString(""),
		ResponseWriter: responseWriter,
	}
}

func (w *GinBodyWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *GinBodyWriter) Bytes() []byte {
	return w.Body.Bytes()
}
