package main

import (
	"fmt"
	"net/http"

	"dexlock.com/todo-project/handlers"
	"github.com/joho/godotenv"

)


func main() {
	err := godotenv.Load(".env")
	if err!=nil{
		fmt.Print("Error loading the .env file")
	}

    fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the route for the login page
	http.HandleFunc("/login", handlers.Login)

	// Define the route for the page that the user will be redirected to after successful login
	http.HandleFunc("/dashboard", handlers.Dashboard)

	http.HandleFunc("/signup", handlers.Signup)

	http.HandleFunc("/renderEditNote", handlers.RenderEditNote)

	http.HandleFunc("/logout", handlers.Logout)

	http.HandleFunc("/createNote", handlers.CreateNote)

	http.HandleFunc("/deleteNote", handlers.DeleteNote)

	// Start the server
	err = http.ListenAndServe(":3000", nil)
	if err!= nil{
		fmt.Print(err.Error())
	}
}

