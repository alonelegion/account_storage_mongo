package router

import (
	v1 "github.com/alonelegion/account_storage_mongo/internal/controllers/v1"
	"github.com/alonelegion/account_storage_mongo/pkg/gincache"
	"github.com/alonelegion/account_storage_mongo/pkg/ginlogger"
	"github.com/alonelegion/account_storage_mongo/pkg/recoverywriter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func NewRouter(logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginlogger.Logger(logger))

	pc := new(v1.PingController)
	router.GET("/", pc.Ping)
	router.GET("/", welcome)
	router.NoRoute(notFound)

	if os.Getenv("SLEEPER") == "true" {
		router.Use(Sleeper())
	}

	// Api
	apiRouter := router.Group("api")
	apiRouter.Use(gin.RecoveryWithWriter(recoverywriter.NewGinRecoverWriter(logger)))
	apiRouter.Use(gincache.CacheResponse(logger))
	mapV1Routes(apiRouter)

	return router
}

func mapV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("v1")
	{
		// Template
		accountsGroup := v1Group.Group("accounts")
		{
			accounts := new(v1.AccountController)
			accountsGroup.GET("/accounts", accounts.GetAllAccounts)
			accountsGroup.POST("/account", accounts.CreateAccount)
			accountsGroup.GET("/account/:accountId", accounts.GetAccount)
			accountsGroup.PUT("/account/:accountId", accounts.EditAccount)
			accountsGroup.DELETE("/account/:accountId", accounts.DeleteAccount)
		}
	}

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
