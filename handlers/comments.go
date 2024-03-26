package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
	"time"
)

func CommentsHanlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllComments(w, r)
	case "POST":
		createComment(w, r)
	case "DELETE":
		deleteComment(w, r)
	}
}

func getAllComments(w http.ResponseWriter, r *http.Request) {
	//parsing
	var commentsData []models.Comments
	byteData,_ := os.ReadFile("db/comments.json")
	json.Unmarshal(byteData, &commentsData)
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
	<h3>Comments list</h3>`
	for i := 0; i < len(commentsData); i++ {
		div := "<div class='list'>"
		Id := "<h4>Comment ID:<i>" + fmt.Sprint(commentsData[i].Id) + "</i></h4>"
		userID := "<h4>User ID:<i>" + fmt.Sprint(commentsData[i].UserId) + "</i></h4>"
		postID := "<h4>Post ID:<i>" + fmt.Sprint(commentsData[i].PostId) + "</i></h4>"
		text := "<h4>Text:<i>" + commentsData[i].Text + "</i></h4>"
		CreatedAt := "<h4>Created at:<i>" + commentsData[i].CreatedAt + "</i></h4>"
		div += Id + userID + postID + text + CreatedAt
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

func createComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comments
	json.NewDecoder(r.Body).Decode(&newComment)


	var commentsData []models.Comments
	byteData,_ := os.ReadFile("db/comments.json")
	json.Unmarshal(byteData, &commentsData)

	for i := 0; i < len(commentsData); i++ {
		if commentsData[i].Id == newComment.Id {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Comment with such kind of ID is already exist")
			return
		}
	}
	newComment.CreatedAt = time.Now().Format(time.RFC1123)
	commentsData = append(commentsData, newComment) 
	
	res,_ := json.Marshal(commentsData)
	os.WriteFile("db/comments.json", res,0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Successfully Created")
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	var deleteComment models.Comments
	json.NewDecoder(r.Body).Decode(&deleteComment)

	var commentsData []models.Comments
	bytedata,_ := os.ReadFile("db/comments.json")
	json.Unmarshal(bytedata, &commentsData)

	for i := 0; i < len(commentsData); i++ {
		if commentsData[i].Id == deleteComment.Id {
			commentsData = append(commentsData[:i],commentsData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "no comment with such kind of ID")
			return
		}
	}

	res,_ := json.Marshal(commentsData)
	os.WriteFile("db/comments.json", res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Successfully Deleted")
}