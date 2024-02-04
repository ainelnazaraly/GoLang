package characters

import (
	"fmt"
	"net/http"
	
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi, my name is Ainel and recently my mind was filled with Hunger Games series so this is my API consisting of the main characters. ")
}