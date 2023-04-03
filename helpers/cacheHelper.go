package helpers

import(
	"net/http"
)

func RemoveCaching(w http.ResponseWriter){
	w.Header().Set("Cache-Control", "no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}