package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
	"time"
)

func RepliesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllReplies(w, r)
	case "POST":
		createReply(w,r)
	case "PUT":
		updateReply(w, r)
	case "DELETE":
		deleteReply(w, r)
	}
}

func getAllReplies(w http.ResponseWriter, r *http.Request) {
	//opening variable for copying data from file (Parsing)
	var repliesData []models.Replies
	// reading json file and saving to variable
	byteData, _ := os.ReadFile("db/replies.json")
	// converting json to variable []models.Replies
	json.Unmarshal(byteData, &repliesData)

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
	<h3>Replies list</h3>` 
	for i := 0; i < len(repliesData); i++ {
		div := "<div class='list'>"
		Id := "<h4>Reply ID:<i>" + fmt.Sprint(repliesData[i].Id) + "</i></h4>"
		userID := "<h4>User ID:<i>" + fmt.Sprint(repliesData[i].UserId) + "</i></h4>"
		postID := "<h4>Post ID:<i>" + fmt.Sprint(repliesData[i].PostId) + "</i></h4>"
		commentID := "<h4>Comment ID:<i>" + fmt.Sprint(repliesData[i].CommentId) + "</i></h4>"
		text := "<h4>Text:<i>" + repliesData[i].Text + "</i></h4>"
		CreatedAt := "<h4>Created at:<i>" + repliesData[i].CreatedAt + "</i></h4>"
		UpdatedAt := "<h4>Updated at:<i>" + repliesData[i].UpdatedAt + "</i></h4>"
		div += Id + userID + postID + commentID + text + CreatedAt + UpdatedAt
		div += "</div>"
		template += div
	}
	template +=  `
	</div>
	</body>
	</html>
	`
	//sending taken result to client
	fmt.Fprint(w, template)
}

func createReply(w http.ResponseWriter, r *http.Request) {
	//parsing r.Body
	var newReply models.Replies
	json.NewDecoder(r.Body).Decode(&newReply)

	//parsing json
	var repliesData [] models.Replies
	byteData, _ := os.ReadFile("db/replies.json")
	json.Unmarshal(byteData, &repliesData)

	for i := 0; i < len(repliesData); i++ {
		if repliesData[i].Id == newReply.Id && repliesData[i].CommentId == newReply.CommentId {	
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "comment with such kind of ID is already exist")	
			return
		}
	}
	newReply.CreatedAt = time.Now().Format(time.RFC1123)
	newReply.UpdatedAt = time.Now().Format(time.RFC1123)
	repliesData = append(repliesData, newReply)
	//array variable to json db file
	res, _ := json.Marshal(repliesData)
	os.WriteFile("db/replies.json",res,0)

	//sendiing response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created Successfully")
}

func updateReply(w http.ResponseWriter, r *http.Request) {
	//parsing r.body from request to variable type models.Replies
	var newReply models.Replies
	json.NewDecoder(r.Body).Decode(&newReply)

	//parsing json 
	var repliesData []models.Replies
	byteData,_ := os.ReadFile("db/replies.json")
	json.Unmarshal(byteData, &repliesData)

	for i := 0; i < len(repliesData); i++ {
		if repliesData[i].Id == newReply.Id && repliesData[i].CommentId == newReply.CommentId {
			newReply.UpdatedAt = time.Now().Format(time.RFC1123)
			repliesData[i].UpdatedAt = newReply.UpdatedAt
			repliesData[i].Text = newReply.Text
		}
	}
	//wrapping gotten data to db file json
	res, _ := json.Marshal(repliesData)
	os.WriteFile("db/replies.json", res,0)

	// sending response
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Updated Succesfully")
}

func deleteReply(w http.ResponseWriter, r *http.Request) {
	var deleteReply models.Replies
	json.NewDecoder(r.Body).Decode(&deleteReply)

	var repliesData []models.Replies
	byteData,_ := os.ReadFile("db/replies.json")
	json.Unmarshal(byteData, &repliesData)

	for i := 0; i < len(repliesData); i++ {
		if repliesData[i].Id == deleteReply.Id {
			repliesData = append(repliesData[:i], repliesData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "no reply with such kind of ID")
		}
	}
	
	res,_ := json.Marshal(repliesData)
	os.WriteFile("db/replies.json", res, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Deleted Successfully")
}