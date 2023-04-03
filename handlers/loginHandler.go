package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"dexlock.com/todo-project/helpers"
)

var store = sessions.NewCookieStore([]byte("random-key"))

// loginHandler renders the login page and handles the form submission
func Login(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")
	helpers.RemoveCaching(w)

	type Response struct {
		Error string `json:"error"`
	}

	type TokenStruct struct {
		Token string `json:token`
	}

	token := helpers.GetToken(w, r)
	if token != "" {

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// Parse the login template
	// tmpl := template.Must(template.ParseFiles("html/login.html"))
	tmpl, err := template.ParseFiles("html/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the request method is POST, the user has submitted the form
	if r.Method == http.MethodPost {
		// Extract the username and password from the form
		credentials, err := json.Marshal(map[string]string{
			"Email":    r.FormValue("Email"),
			"Password": r.FormValue("Password"),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send the credentials to the login API

		resp, err := http.Post(baseUrl+":8080/users/login", "application/json", bytes.NewBuffer(credentials))


		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// If the API returns a status code indicating success (e.g. 200 or 204), redirect the user to the dashboard
		if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
			body, _ := ioutil.ReadAll(resp.Body)
			var tokenStruct TokenStruct
			json.Unmarshal(body, &tokenStruct)

			session, err := store.Get(r, "login-session")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			session.Values["token"] = tokenStruct.Token
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}


			req, err := http.NewRequest(http.MethodPatch, baseUrl+":8080/user/active-status", nil)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			req.Header.Set("Token", tokenStruct.Token)
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
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

		defer resp.Body.Close()
		return

	}

	// If the request method is not POST, render the login page with an empty error message
	tmpl.Execute(w, nil)
}
