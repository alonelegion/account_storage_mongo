package routes

import (
	"github.com/alonelegion/account_storage_mongo/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/accounts", controllers.GetAllAccounts)
	router.POST("/account", controllers.CreateAccount)
	router.GET("/account/:accountId", controllers.GetAccount)
	router.PUT("/account/:accountId", controllers.EditAccount)
	router.DELETE("/account/:accountId", controllers.DeleteAccount)
	router.NoRoute(notFound)
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
