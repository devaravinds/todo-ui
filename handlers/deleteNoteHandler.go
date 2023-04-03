package handlers

import (
	"fmt"
	"net/http"
	"os"

	"dexlock.com/todo-project/helpers"
)

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")
	helpers.RemoveCaching(w)
	id := r.URL.Query().Get("id")

	req, err := http.NewRequest(http.MethodDelete, baseUrl+":8080/note/"+id, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
}
