package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"
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
	w.Header().Set("Content-T ype", "application/json")
	listUser = nil
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

func GetAllUser() []Person {
	var result []Person = nil
	collection := client.Database("mydb").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		result = append(result, person)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}
	return result
}
func GetPeopleEndPoint1(w http.ResponseWriter, r *http.Request) {

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

func SignInEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Functions is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	listUser = nil
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
	if listUser != nil {
		for _, item := range listUser {
			if item.Email == params["email"] && item.Password == params["password"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	} else {
		json.NewEncoder(w).Encode("Don't have account")
		return
	}
	json.NewEncoder(w).Encode("Doo't have account")
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func createIdUser(collection *mongo.Collection, ctx context.Context, email string) string {
	cursor, err := collection.Find(ctx, bson.M{})
	var people []Person
	if err != nil {
		return "user 1"
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
		if person.Email == email {
			return "error"
		}
	}
	if err := cursor.Err(); err != nil {
		return "user 1"
	}
	result := "user " + strconv.Itoa(len(people)+1)
	return result
}

func checkFound(listUser []Person, email string) bool {
	for _, item := range listUser {
		if item.Email == email {
			return false
		}
	}
	return true
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUser function is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//collection := client.Database("mydb").Collection("user")
	//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var listUser []Person = GetAllUser()
	if params["email"] != "" {
		if !checkFound(listUser, params["email"]) {

		} else {
			json.NewEncoder(w).Encode("Error email not found")
		}
	} else {
		json.NewEncoder(w).Encode("Error email is null")
		return
	}
}
func SignUpUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create user is called")
	w.Header().Set("Content-Type", "application/json")
	var person Person
	params := mux.Vars(r)
	collection := client.Database("mydb").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if params["email"] != "" && params["password"] != "" && params["name"] != "" && params["age"] != "" {
		if !valid(params["email"]) || len(params["password"]) < 7 {
			json.NewEncoder(w).Encode("Email or password is not format")
			return

		} else {
			person.Email = params["email"]
			person.Name = params["name"]
			person.Password = params["password"]
			i, _ := strconv.Atoi(params["age"])
			person.Age = i
			person.Id = createIdUser(collection, ctx, person.Email)
			if person.Id == "error" {
				json.NewEncoder(w).Encode("User already exists")
				return
			}
		}
	} else {
		json.NewEncoder(w).Encode("Field is not null")
		return
	}
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(w).Encode(result)
}

func main() {
	fmt.Println("Starting the application.....")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()

	//Handle Function
	router.HandleFunc("/user/createuser", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/user/getAll", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/find/{id}", GetPersonFromDatabaseEndPoint).Methods("GET")
	router.HandleFunc("/user/signIn/{email}/{password}", SignInEndPoint).Methods("GET")
	router.HandleFunc("/user/signUp/{email}/{password}/{name}/{age}", SignUpUserEndPoint).Methods("POST")
	router.HandleFunc("/api/user/deleteUser/{email}", DeleteUserEndPoint).Methods("DELETE")

	http.ListenAndServe(":2011", router)
}
