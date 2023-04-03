package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"dexlock.com/todo-project/helpers"
)

type Note struct {
	ID        string `json:"ID"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	User_id   string `json:"user_id"`
	Completed bool   `json:"completed"`
}

type ActiveUsers struct {
	Count string `json:"count"`
}

// dashboardHandler renders the dashboard page
func Dashboard(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")
	helpers.RemoveCaching(w)

	// Parse the dashboard template
	tmpl := template.Must(template.ParseFiles("html/dashboard.html"))

	// Make a GET request to retrieve notes
	req, err := http.NewRequest("GET", baseUrl+":8080/notes", nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	body, err := helpers.GetBody(req, w, r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var notes []Note
	err = json.Unmarshal(body, &notes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userCountReq, err := http.NewRequest("GET", baseUrl+":8080/active-users", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	body, err = helpers.GetBody(userCountReq, w, r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var activeUsers ActiveUsers

	activeUsers.Count = string(body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	data := map[string]interface{}{
		"Notes":       notes,
		"ActiveUsers": activeUsers.Count,
	}

	// Create a map of data to pass to the template

	tmpl.Execute(w, data)

}
