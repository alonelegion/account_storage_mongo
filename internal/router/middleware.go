package router

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Sleeper() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(3)
	}
}
