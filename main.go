package main

import (
	"fmt"
	"instagram/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.GetHomePage)

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/comments", handlers.CommentsHanlder)
	http.HandleFunc("/replies", handlers.RepliesHandler)

	fmt.Println("Server working on port :1000")
	http.ListenAndServe(":1000", nil)
}
