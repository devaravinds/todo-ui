package helpers

import (
	"net/http"
	"time"
	"github.com/gorilla/sessions"
)

func GetToken(w http.ResponseWriter, r *http.Request) string{
	var store = sessions.NewCookieStore([]byte("random-key"))
	session, err := store.Get(r, "login-session")
	if err!=nil{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	if token, ok := session.Values["token"].(string); ok {
		return token
	} else {
		return ""
	}
}

func DeleteToken(w http.ResponseWriter, r *http.Request){
	cookies := r.Cookies()
	for _, cookie := range cookies {
		cookie := &http.Cookie{
			Name:    cookie.Name,
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(w, cookie)
	}
	
}