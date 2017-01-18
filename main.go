//main.go

package main

import (
	"net/http"
	"log"
)

func main() {
	const PORT = ":5050"

	log.Printf("goJWT-auth starting on port: %s", PORT)

	http.HandleFunc("/auth/register", notYetImplemented)
	http.HandleFunc("/auth/login", notYetImplemented)
	http.HandleFunc("/auth/logout", notYetImplemented)
	http.HandleFunc("/auth/verify", notYetImplemented)
	http.HandleFunc("/auth/", fourOhFour)
	http.ListenAndServe(PORT, nil)	
}

func notYetImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(210)
	w.Write([]byte("Not Yet Implemented"))
	return
}

func fourOhFour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	return
}