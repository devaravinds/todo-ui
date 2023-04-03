package helpers

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func GetBody(req *http.Request, w http.ResponseWriter, r *http.Request) ([]byte,error){

	token := GetToken(w,r)
	req.Header.Set("Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	if resp.StatusCode == http.StatusUnauthorized{
		http.Redirect(w, r, "/login",http.StatusSeeOther)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	return body,err
}