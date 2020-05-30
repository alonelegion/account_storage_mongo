package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/api/accounts", getAccounts).Methods("GET")
	router.GET("/api/accounts/{id}", getAccount).Methods("GET")
	router.HandleFunc("/api/accounts", createAccount).Methods("POST")
	router.HandleFunc("/api/accounts/{id}", updateAccount).Methods("PUT")
	router.HandleFunc("/api/accounts/{id}", deleteAccount).Methods("DELETE")
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
