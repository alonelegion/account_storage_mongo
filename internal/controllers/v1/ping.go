package v1

import (
	"github.com/alonelegion/account_storage_mongo/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct{}

func (pc *PingController) Ping(c *gin.Context) {

	if err := application.DB().DB().Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "unavailable",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

	return
}
