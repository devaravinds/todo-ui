package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")

	type Response struct {
		Error string `json:"error"`
	}

	tmpl := template.Must(template.ParseFiles("html/signup.html"))

	if r.Method == http.MethodPost {
		userData, err := json.Marshal(map[string]string{
			"First_Name": r.FormValue("First Name"),
			"Last_Name":  r.FormValue("Last Name"),
			"Email":      r.FormValue("Email"),
			"Password":   r.FormValue("Password"),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		resp, err := http.Post(baseUrl+":8080/users/signup", "application.json", bytes.NewBuffer(userData))


		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if resp.StatusCode != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var response Response
			json.Unmarshal(body, &response)
			tmpl.Execute(w, response)
		}

		return
	}

	tmpl.Execute(w, nil)
}
