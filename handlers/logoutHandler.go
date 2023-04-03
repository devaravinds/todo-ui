package handlers

import (
	"fmt"
	"net/http"
	"os"

	"dexlock.com/todo-project/helpers"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	baseUrl := os.Getenv("BASE_URL")
	req, err := http.NewRequest(http.MethodPatch, baseUrl+":8080/user/active-status", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.Header.Set("Token", helpers.GetToken(w, r))
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	helpers.DeleteToken(w, r)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
