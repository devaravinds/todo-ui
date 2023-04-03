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

func RenderEditNote(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")
	tmpl, err := template.ParseFiles("html/editNote.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("id")

	noteData, err := json.Marshal(map[string]string{
		"title":   r.FormValue("Title"),
		"content": r.FormValue("Content"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}



	req, err := http.NewRequest(http.MethodPatch, baseUrl+":8080/note/"+id, bytes.NewBuffer(noteData))
	if err!=nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	req.Header.Set("Token", helpers.GetToken(w, r))

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req, err = http.NewRequest("GET", baseUrl+":8080/note/"+id, nil)



	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	body, err := helpers.GetBody(req, w, r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var note Note

	err = json.Unmarshal(body, &note)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = tmpl.Execute(w, note)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
