package main

import (
	"fmt"
	"os"

	"github.com/smstong/basicauth"
)

func main() {
	auth := basicauth.NewBasicAuth()
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s username password", os.Args[0])
		return
	}
	username := os.Args[1]
	password := os.Args[2]

	fmt.Printf("%s:%s\n", username, auth.EncryptPass(password))
}
