package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
	"time"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("users ishladi")
	switch r.Method {
	case "GET":
		getAllUsers(w, r)
	case "POST":
		createUser(w, r)
	case "PUT":
		updateUser(w, r)
	case "DELETE":
		deleteUser(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// req.Body parse qilamiz  | User
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// File ni ochamiz, parse qilamiz | []User
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userdata)

	// yangi userni arrayga qo'shamiz

	for i := 0; i < len(userdata); i++ {
		if userdata[i].Username == newUser.Username {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "This username is already taken by someone")
			return
		}
	}

	newUser.Id = len(userdata)+1
	newUser.CreatedAt = time.Now().Format(time.RFC1123)
	newUser.UpdatedAt = time.Now().Format(time.RFC1123)
	userdata = append(userdata, newUser)

	// array ni faylga yozamiz
	res, _ := json.Marshal(userdata)
	os.WriteFile("db/users.json",res,0)

	// yangi userni jsonga o'zgaritirb responsega yozamiz | 
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userdata)

	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>doc</title>
		<style>
		body {
			margin: 0;
			padding: 0;
			background-color: whitesmoke;
		}
		.container {
			text-align: center;
			margin: 0px;
			padding: 10px;
		}
		.list {
			text-align: left;
			max-width: 400px;
			border: 1px solid grey;
			border-radius: 4px;
			margin: 20px;
			padding: 10px;
			background-color: lightblue;
			font-family: sans-serif;
		}
		.list h1,h2,h3,h4,h5,h6 {
			margin-right: 5px;
			color: grey;
			margin: 10px;
			padding: 4px;
			border: 1px solid grey;
			border-radius: 4px;
		}
		.list i {
			color: black;
			font-family: calibri;
			font-size: 16px;
		}
		</style>
	</head>
	<body>
	<div class='container'>
	<a href='/'><button>Home</button></a>
	<h3>Users list</h3>` 
	for i := 0; i < len(userdata); i++ {
		div := "<div class='list'>"
		Id := "<h4>User ID:<i>" + fmt.Sprint(userdata[i].Id) + "</i></h4>"
		username := "<h4>Username:<i>" + userdata[i].Username + "</i></h4>"
		Email := "<h4>Email:<i>" + userdata[i].Email + "</i></h4>"
		Age := "<h4>Age:<i>" + fmt.Sprint(userdata[i].Age) + "</i></h4>"
		CreatedAt := "<h4>Created at:<i>" + userdata[i].CreatedAt + "</i></h4>"
		UpdatedAt := "<h4>Updated at:<i>" + userdata[i].UpdatedAt + "</i></h4>"
		div += Id + username + Email + Age + CreatedAt + UpdatedAt
		div += "</div>"
		template += div
	}
	template +=  `
	</div>
	</body>
	</html>
	`
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, template)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// req.Body parse qilamiz  | User
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// File ni ochamiz, parse qilamiz | []User
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
 	json.Unmarshal(byteData, &userdata)

	// yangi userni arrayga qo'shamiz
	for i := 0; i < len(userdata); i++ {
		if userdata[i].Id == newUser.Id {
			userdata[i].Username = newUser.Username
			userdata[i].Email = newUser.Email
			userdata[i].Age = newUser.Age
			userdata[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}

	// array ni faylga yozamiz
	res, _ := json.Marshal(userdata)
	os.WriteFile("db/users.json",res,0)

	// yangi userni jsonga o'zgaritirb responsega yozamiz | 
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(newUser)
	fmt.Fprint(w,"Updated Succesfully")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//parsing
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)

	// file to array variable
	var userdata []models.User
	byteData, _ := os.ReadFile("db/users.json")
	json.Unmarshal(byteData, &userdata)

	for i := 0; i < len(userdata); i++ {
		if userdata[i].Id == newUser.Id {
			userdata = append(userdata[:i], userdata[i+1:]...)
		}
	}

	res, _ := json.Marshal(userdata)
	os.WriteFile("db/users.json",res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Deleted Succesfully")
}
