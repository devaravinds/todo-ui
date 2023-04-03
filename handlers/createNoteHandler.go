package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"dexlock.com/todo-project/helpers"
)



func CreateNote(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")
	tmpl, err := template.ParseFiles("html/editNote.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	noteData, err := json.Marshal(map[string]string{
		"title":   r.FormValue("Title"),
		"content": r.FormValue("Content"),
	})

	req, err := http.NewRequest(http.MethodPost, baseUrl+":8080/notes", bytes.NewBuffer(noteData))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Token", helpers.GetToken(w, r))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode == http.StatusOK {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

	tmpl.Execute(w, nil)
}
