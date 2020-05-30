package main

import (
	"context"
	"encoding/json"
	"github.com/alonelegion/account_storage_mongo/config"
	"github.com/alonelegion/account_storage_mongo/helper"
	"github.com/alonelegion/account_storage_mongo/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func main() {

	// Database
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes

	route := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", route))
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var account models.AccountInfo

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&account)

	update := bson.D{
		{"$set", bson.D{
			{"title", account.Title},
			{"account_auth", bson.D{
				{"login", account.AccountAuth.Login},
				{"password", account.AccountAuth.Password},
			}},
			{"email", account.Email},
			{"phone", account.Phone},
			{"reserve_email", account.ReserveEmail},
			{"owner", account.Owner},
			{"notice", account.Notice},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&account)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	account.ID = id

	json.NewEncoder(w).Encode(account)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account models.AccountInfo

	_ = json.NewDecoder(r.Body).Decode(&account)

	collection := helper.ConnectDB()

	result, err := collection.InsertOne(context.TODO(), account)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account models.AccountInfo

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&account)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var accounts []models.AccountInfo

	collection := helper.ConnectDB()

	curl, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer curl.Close(context.TODO())

	for curl.Next(context.TODO()) {
		var account models.AccountInfo

		err := curl.Decode(&account)
		if err != nil {
			log.Fatal(err)
		}

		accounts = append(accounts, account)
	}

	if err := curl.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(accounts)
}
