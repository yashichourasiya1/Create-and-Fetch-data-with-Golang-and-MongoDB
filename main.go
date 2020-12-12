//siddharth-star/create-and-fetch-data-with-goland-and-mongodb
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"context"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
//    "go.mongodb.org/mongo-driver/mongo/readpref"
)

// Book struct (Model)go.mongodb.org/mongo-driver/mongo
type User struct {
	ID     string  `json:"id",omitempty`
	Name   string  `json:"Name",omitempty`
	Date_Of_Birth  string  `json:"DOB",omitempty`
	Phone_Number   string  `json:"Phone",omitempty`
	Email_Address string  `json:"Email",omitempty`
	Time_stamp string  `json:"Time_stamp",omitempty`

}

type Contact struct {
	UserIdOne  string  `json:"UserIdOne",omitempty`
	UserIdTwo  string  `json:"UserIdTwo",omitempty`
	Time_stamp string  `json:"Time_stamp",omitempty`

}
// Init users var as a slice user , contact struct
var users []User
var contacts []Contact
//get all users 
func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
// Fetch all contacts 
func getCon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}
// Get single user with id
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through users and find one with the id from the params
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
// Check for 14 days 
func get14(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through users and find one with the id from the params
	fmt.Println("“=========time stamps recogination Started==========”")
	// log.Println(params["id"],params[])
	for _, item := range contacts {
		if (item.UserIdOne == params["id"]) || (item.UserIdTwo == params["id"]) {
			v1,_:=strconv.ParseInt(params["infection_timestamp"],0,64)
			v2,_:=strconv.ParseInt(item.Time_stamp,0,64)
			fmt.Println("“time stamps v1 and v2”",v1,v2)
			if v1-v2<15000000{
				json.NewEncoder(w).Encode(item)
			
		}
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
// Add new User 
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book User
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe "mongodb://localhost:27017"
	book.Time_stamp = time.Now().Format("20060102150405")
	users = append(users, book)
	json.NewEncoder(w).Encode(book)
	//=============
	var Client *mongo.Client
 
// ConnectDatabase is used to connect the MongoDB database

    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
 
    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    Client = client
    if err != nil {
        log.Fatal(err)
    }
 
    // Check the connection
    err = Client.Ping(context.TODO(), nil)
 
    if err != nil {
        log.Fatal(err)
    }
 
	log.Println("Database Connected.")
	//==========
	collection := client.Database("appointy").Collection("User")

insertResult, err := collection.InsertOne(context.TODO(), book)

if err != nil {
log.Fatal(err)
}
fmt.Println("“Inserted post with ID:”", insertResult.InsertedID)

}

func createCon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Contact
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Time_stamp = time.Now().Format("20060102150405") // Mock ID - not safe "mongodb://localhost:27017"
	contacts = append(contacts, book)
	json.NewEncoder(w).Encode(book)
	//=============
	var Client *mongo.Client
 
// ConnectDatabase is used to connect the MongoDB database

    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
 
    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)
    Client = client
    if err != nil {
        log.Fatal(err)
    }
 
    // Check the connection
    err = Client.Ping(context.TODO(), nil)
 
    if err != nil {
        log.Fatal(err)
    }
 
	log.Println("Database Connected.")
	
	collection := client.Database("appointy").Collection("Contacts")

insertResult, err := collection.InsertOne(context.TODO(), book)

if err != nil {
log.Fatal(err)
}
fmt.Println("“Inserted post with ID:”", insertResult.InsertedID)

}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()


	users = append(users, User{ID: "2", Name: "454555", Date_Of_Birth: "12/12/2007",Phone_Number:"86764" , Email_Address:"vfdfg"})

	// Route handles & endpoints
	r.HandleFunc("/users", getusers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/contacts", createCon).Methods("POST")
	r.HandleFunc("/contacts", getCon).Methods("GET")
	r.HandleFunc("/contacts/{id}&{Time_stamp}", get14).Methods("GET")
	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
