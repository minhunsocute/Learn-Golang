package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json: "age"`
	Id       string `json:"id"`
}

var client *mongo.Client
var listUser []Person

func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Function is called")
	w.Header().Set("Content-Type", "application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	collection := client.Database("mydb").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(w).Encode(result)
}
func GetPeopleEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Functions get is called")
	w.Header().Set("Content-Type", "application/json")
	collection := client.Database("mydb").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		listUser = append(listUser, person)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(listUser)
}

func GetPersonEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Find user is called")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range listUser {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPersonFromDatabaseEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fin user from database is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	var person Person
	collection := client.Database("mydb").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{Id: params["id"]}).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(person)
}
func main() {
	fmt.Println("Starting the application.....")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()

	//Handle Function
	router.HandleFunc("/createuser", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/getAll", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/find/{id}", GetPersonFromDatabaseEndPoint).Methods("GET")

	http.ListenAndServe(":2011", router)
}
