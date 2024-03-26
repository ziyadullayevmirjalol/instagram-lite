package handlers

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, 
	`<!DOCTYPE html>
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
	<h3>Home page</h3> 
		<div class='list'>
		<a href='/posts'><button>All posts</button></a>
		<a href='/comments'><button>All comments</button></a>
		<a href='/replies'><button>All replies</button></a>
		<a href='/users'><button>All users</button></a>
		</div>
	</div>
	</body>
	</html>
	`)
}
