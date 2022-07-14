package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
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

type User struct {
	Id          string `json:"id"`
	Name        string `json:name`
	Email       string `json:email`
	Password    string `json:password`
	Avatar      string `json:avatar`
	PhoneNumber string `json:phoneNumber`
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

func remove(list []Person, email string) []Person {
	var result []Person = nil
	for _, item := range list {
		if item.Email != email {
			result = append(result, item)
		}
	}
	return result
}

func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUser function is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mydb").Collection("user")
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var listUser []Person = GetAllUser()
	if params["email"] != "" {
		if !checkFound(listUser, params["email"]) {
			deleteResult, _ := collection.DeleteOne(context.TODO(), bson.M{"email": params["email"]})
			if deleteResult.DeletedCount == 0 {
				json.NewEncoder(w).Encode("Error cannot delete User")
				return
			}
			listUser = remove(listUser, params["email"])
			json.NewEncoder(w).Encode(listUser)
			return
		} else {
			json.NewEncoder(w).Encode("Error email not found")
			return
		}
	} else {
		json.NewEncoder(w).Encode("Error email is null")
		return
	}
}

func EditUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EditUser function is called")
	w.Header().Set("Content-Type", "application/json")

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

func postUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post user is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mydb").Collection("users")
	var person Person
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if params["email"] != "" && params["password"] != "" && params["name"] != "" && params["age"] != "" {
		if !valid(params["email"]) || len(params["password"]) < 7 {
			json.NewEncoder(w).Encode("Error . Email is not format")
			return
		} else {
			person.Email = params["email"]
			person.Name = params["name"]
			person.Password = params["password"]
			i, _ := strconv.Atoi(params["age"])
			person.Age = i
			person.Id = createIdUser(collection, ctx, person.Email)
			if person.Id == "error" {
				json.NewEncoder(w).Encode("Error . User is already")
				return
			}

		}
	} else {
		json.NewEncoder(w).Encode("Email . Field is not null")
		return
	}
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(w).Encode(result)
	return
}
func editUserEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit user is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mydb").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if params["email"] != "" && params["password"] != "" && params["name"] != "" && params["age"] != "" {
		if !valid(params["email"]) || len("password") < 7 {
			json.NewEncoder(w).Encode("Error . Email is not Format")
			return
		} else {
			//var person Person
			// err := collection.FindOne(ctx, Person{Email: params["email"]}).Decode(&person)
			// fmt.Println(params["email"])
			// fmt.Println(person.Email)
			// fmt.Println(err)
			// if err != nil {
			// 	json.NewEncoder(w).Encode("Error . User is not found")
			// 	return
			// }
			result, err := collection.ReplaceOne(ctx, bson.M{"email": params["email"]}, bson.M{
				"name":     params["name"],
				"email":    params["email"],
				"password": params["password"],
				"age":      params["age"],
				"id":       params["id"],
			})

			if err != nil {
				json.NewEncoder(w).Encode("Error . Edit user is fail")
				return
			}
			json.NewEncoder(w).Encode(result)
			return
		}
	} else {
		json.NewEncoder(w).Encode("Error . Field is null")
		return
	}
	return
}

//-------------------------------------------Mountain trip api

func get_allUser() []User {
	var result []User = nil
	collection := client.Database("mountain_trip").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person User
		cursor.Decode(&person)
		result = append(result, person)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}
	return result
}

var list_user []User

func SignInUser_EndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign In User func is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params["email"] != "" && params["password"] != "" {
		if !valid(params["email"]) || len(params["password"]) < 7 {
			json.NewEncoder(w).Encode("Error Email or pasword is not format")
			return
		} else {
			list_user = get_allUser()
			for _, item := range list_user {
				fmt.Println(item.Email + " " + item.Password)
				if item.Email == params["email"] && item.Password == params["password"] {
					json.NewEncoder(w).Encode(item)
					return
				}
			}
			json.NewEncoder(w).Encode("Error User is not found")
			return
		}
	} else {
		json.NewEncoder(w).Encode("Error Email,Password is not null")
		return
	}
}

func SignUpUser_EndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign Up User func in called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mountain_trip").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if params["email"] != "" && params["password"] != "" && params["phone"] != "" {
		if !valid(params["email"]) || len(params["password"]) < 7 {
			json.NewEncoder(w).Encode("Error")
			return
		} else {
			list_user = nil
			list_user = get_allUser()
			for _, item := range list_user {
				if item.Email == params["email"] {
					json.NewEncoder(w).Encode("Error User already exists")
					return
				}
			}
			var user User
			index := strings.Index(params["email"], "@")
			name := params["email"][0:index]

			id := "User " + strconv.Itoa(len(list_user))

			user.Id = id
			user.Name = name
			user.Email = params["email"]
			user.Password = params["password"]
			user.PhoneNumber = params["phone"]
			user.Avatar = "null"

			result, _ := collection.InsertOne(ctx, user)
			fmt.Println(result)
			json.NewEncoder(w).Encode(user)
			return

		}
	} else {
		json.NewEncoder(w).Encode("Error Field is not null")
		return
	}
}

func EditUser_EndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit Func is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mountain_trip").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if params["name"] != "" && params["password"] != "" && params["phone"] != "" {
		if !valid(params["email"]) || len(params["password"]) < 7 {
			json.NewEncoder(w).Encode("Error Email or password is invalid")
			return
		} else {
			var user User
			filter := bson.D{{"id", params["id"]}}
			err := collection.FindOne(ctx, filter).Decode(&user)
			if err != nil {
				json.NewEncoder(w).Encode("Error User is not found")
				return
			}
			update := bson.D{{"$set", bson.D{{"name", params["name"]},
				{"email", params["email"]},
				{"phoneNumber", params["phone"]}}}}
			result, err := collection.UpdateOne(context.TODO(), filter, update)
			// err = collection.Update(bson.M{"_id": user.ID},bson.M{"$set": bson.M{"name": params["name"]}})
			if err != nil {
				json.NewEncoder(w).Encode("Error Can not update Profile User")
				return
			}
			json.NewEncoder(w).Encode(result)

			return
		}
	} else {
		json.NewEncoder(w).Encode("Error Input is null")
		return
	}
}

func Reset_EndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Forgot passwordd func is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mountain_trip").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if len(params["newPass"]) >= 7 {
		filter := bson.D{{"email", params["email"]}}
		update := bson.D{{"$set", bson.D{{"password", params["newPass"]}}}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			json.NewEncoder(w).Encode("Error Cann't Change your password")
			return
		}

		var user User
		err = collection.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			json.NewEncoder(w).Encode("Error user is not found")
			return
		}
		json.NewEncoder(w).Encode(user)
		return
	} else {
		json.NewEncoder(w).Encode("Error Password must be more than 7 characters")
		return
	}
}

func ChangePassword_EndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Change Password func is called")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	collection := client.Database("mountain_trip").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if len(params["newPass"]) >= 7 {
		filter := bson.D{{"id", params["id"]}}
		var user User
		err := collection.FindOne(ctx, filter).Decode(&user)

		fmt.Println(user.Password + " - " + params["yourPass"])

		if user.Password != params["yourPass"] {
			json.NewEncoder(w).Encode("Error Your Password is invalid")
			return
		}

		update := bson.D{{"$set", bson.D{{"password", params["newPass"]}}}}
		result, err := collection.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			json.NewEncoder(w).Encode("Error Cann't Change your password")
			return
		}
		json.NewEncoder(w).Encode(result)
		return
	} else {
		json.NewEncoder(w).Encode("Error Password must be more than 7 characters")
		return
	}
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
	//router.HandleFunc("/user/signIn/{email}/{password}", SignInEndPoint).Methods("GET")
	//router.HandleFunc("/user/signUp/{email}/{password}/{name}/{age}", SignUpUserEndPoint).Methods("POST")
	router.HandleFunc("/user/deleteUser/{email}", DeleteUserEndPoint).Methods("DELETE")
	router.HandleFunc("/user/editUser/{email}/{password}/{name}/{age}/{id}", editUserEndPoint).Methods("PATCH")

	// Handle User model
	router.HandleFunc("/user/signIn/{email}/{password}", SignInUser_EndPoint).Methods("GET")
	router.HandleFunc("/user/signUp/{email}/{password}/{phone}", SignUpUser_EndPoint).Methods("POST")
	router.HandleFunc("/user/edit/{id}/{name}/{email}/{password}/{avatar}/{phone}", EditUser_EndPoint).Methods("PATCH")
	router.HandleFunc("/user/changePass/{id}/{newPass}/{yourPass}", ChangePassword_EndPoint).Methods("PATCH")
	router.HandleFunc("/user/resetPass/{email}/{newPass}", Reset_EndPoint).Methods("PATCH")
	// Hanedle place Model

	http.ListenAndServe(":2011", router)
}
