package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	getuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type AccountController struct{}

type Account struct {
	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	Title        string `json:"title,omitempty" bson:"title,omitempty"`
	Login        string `json:"login,omitempty" bson:"login,omitempty"`
	Password     string `json:"password,omitempty" bson:"password,omitempty"`
	Site         string `json:"site,omitempty" bson:"site,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	Phone        int    `json:"phone,omitempty" bson:"phone,omitempty"`
	ReserveEmail string `json:"reserve_email,omitempty" bson:"reserve_email,omitempty"`
	Owner        string `json:"owner,omitempty" bson:"owner,omitempty"`
	Notice       string `json:"notice,omitempty" bson:"notice,omitempty"`
}

// Database instance
var collection *mongo.Collection

func AccountCollection(c *mongo.Database) {
	collection = c.Collection("accounts")
}

func (ac AccountController) GetAllAccounts(c *gin.Context) {
	accounts := []Account{}

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Ошибка при получение всех аккаунтов, Причина: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,

			"message": "Что-то пошло не так!",
		})

		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var account Account
		cursor.Decode(&account)
		accounts = append(accounts, account)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,

		"message": "Все Аккаунты",

		"data": accounts,
	})

	return
}

func (ac AccountController) CreateAccount(c *gin.Context) {
	var account Account

	c.BindJSON(&account)

	id := getuid.New().String()
	title := account.Title
	login := account.Login
	password := account.Password
	site := account.Site
	email := account.Email
	phone := account.Phone
	reserve_email := account.ReserveEmail
	owner := account.Owner
	notice := account.Notice
	newAccount := Account{
		ID:           id,
		Title:        title,
		Login:        login,
		Password:     password,
		Site:         site,
		Email:        email,
		Phone:        phone,
		ReserveEmail: reserve_email,
		Owner:        owner,
		Notice:       notice,
	}

	_, err := collection.InsertOne(context.TODO(), newAccount)
	if err != nil {
		log.Printf("Ошибка при добавлении нового аккаунта в базу данных. Причина: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,

			"message": "Что-то пошло не так",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,

		"message": "Аккаунт успешно создан",
	})

	return
}

func (ac AccountController) GetAccount(c *gin.Context) {
	accountId := c.Param("accountId")

	account := Account{}

	err := collection.FindOne(context.TODO(), bson.M{"id": accountId}).Decode(&account)

	if err != nil {
		log.Printf("Ошибка при получение аккаунта. Причина: %v\n", err)

		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,

			"message": "Аккаунт не найден",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,

		"message": "Один аккаунт",

		"data": account,
	})

	return
}

func (ac AccountController) EditAccount(c *gin.Context) {
	accountId := c.Param("accountId")
	var account Account
	c.BindJSON(&account)

	title := account.Title
	login := account.Login
	password := account.Password
	email := account.Email
	phone := account.Phone
	reserve_email := account.ReserveEmail
	owner := account.Owner
	notice := account.Notice

	newData := bson.M{
		"$set": bson.M{
			"title":         title,
			"login":         login,
			"password":      password,
			"email":         email,
			"phone":         phone,
			"reserve_email": reserve_email,
			"owner":         owner,
			"notice":        notice,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": accountId}, newData)

	if err != nil {
		log.Printf("Ошибка, Причина: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,

			"message": "Что-то пошло не так",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,

		"message": "Аккаунт был изменен",
	})

	return
}

func (ac AccountController) DeleteAccount(c *gin.Context) {
	accountId := c.Param("accountId")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": accountId})

	if err != nil {
		log.Printf("Ошибка при удалении аккаунта. Причина: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,

			"message": "Что-то пошло не так",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,

		"message": "Аккаунт успешно удален",
	})

	return
}
