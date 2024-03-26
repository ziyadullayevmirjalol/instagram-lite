package handlers

import (
	"encoding/json"
	"time"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllPosts(w, r)
	case "POST":
		createPost(w, r)
	case "PUT":
		updatePost(w, r)
	case "DELETE":
		deletePost(w, r)
	}
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	var postsData []models.Posts
	byteData,_ := os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)


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
	<h3>Posts</h3>` 
	for i := 0; i < len(postsData); i++ {
		div := "<div class='list'>"
		Id := "<h4>Post ID:<i>" + fmt.Sprint(postsData[i].Id) + "</i></h4>"
		userID := "<h4>User ID:<i>" + fmt.Sprint(postsData[i].UserId) + "</i></h4>"
		Title := "<h4>Title:<i>" + postsData[i].Title + "</i></h4>"
		Content := "<h4>Content:<i>" + postsData[i].Content + "</i></h4>"
		LikesCount := "<h4>Likes count:<i>" + fmt.Sprint(postsData[i].LikesCount) + "</i></h4>"
		CreatedAt := "<h4>Created at:<i>" + postsData[i].CreatedAt + "</i></h4>"
		UpdatedAt := "<h4>Updated at:<i>" + postsData[i].UpdatedAt + "</i></h4>"
		div += Id + userID + Title + Content + LikesCount + CreatedAt + UpdatedAt
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
func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Posts
	json.NewDecoder(r.Body).Decode(&newPost)

	var postsData []models.Posts
	byteData,_:= os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	for i := 0; i < len(postsData); i++ {
		if postsData[i].Id == newPost.Id {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Post with such kind of ID already exist")
			return
		} 
	}
	
	newPost.Id = len(postsData)+1
	newPost.CreatedAt = time.Now().Format(time.RFC1123)
	newPost.UpdatedAt = time.Now().Format(time.RFC1123)
	postsData = append(postsData, newPost)

	res,_ := json.Marshal(postsData) 
	os.WriteFile("db/posts.json",res,0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>doc</title>
		<link rel="stylesheet" href="assets/css/register.css">
	</head>
	<body>
	<div>
	<h1>Created Successfully</h1>
	</div>
	</body>
	</html>
	`)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost models.Posts
	json.NewDecoder(r.Body).Decode(&updatePost)

	var postsData []models.Posts
	byteData, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	for i := 0; i < len(postsData); i++ {
		if postsData[i].Id == updatePost.Id {
			postsData[i].UserId = updatePost.UserId
			postsData[i].Title = updatePost.Title
			postsData[i].Content = updatePost.Content
			postsData[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}

	res, _ := json.Marshal(postsData)
	os.WriteFile("db/posts.json",res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>doc</title>
		<link rel="stylesheet" href="assets/css/register.css">
	</head>
	<body>
	<div>
	<h1>Updated Successfully</h1>
	</div>
	</body>
	</html>
	`)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	var deletePost models.Posts
	json.NewDecoder(r.Body).Decode(&deletePost)

	var postsData []models.Posts
	byteData, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(byteData, &postsData)

	for i := 0; i < len(postsData); i++ {
		if postsData[i].Id == deletePost.Id {
			postsData = append(postsData[:i],postsData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w,"no posts with such kind of ID")
			return
		}
	}

	res, _ := json.Marshal(postsData)
	os.WriteFile("db/posts.json",res,0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>doc</title>
    <link rel="stylesheet" href="assets/css/register.css">
</head>
<body>
<div>
<h1>Deleted Successfully</h1>
</div>
</body>
</html>
`)
}