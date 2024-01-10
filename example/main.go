package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smstong/basicauth"
)

func main() {
	auth := basicauth.NewBasicAuth()
	if err := auth.LoadUsersFromFile("userdb"); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this page is not procted.")
	})
	http.HandleFunc("/protected", auth.Auth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this page is protected by basic auth.")
	}))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
