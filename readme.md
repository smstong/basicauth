# HTTP Basic Authn for Golang
## TLDR;
```
// create a BasicAuth object
auth := basicauth.NewBasicAuth()

// add user/password 
auth.AddUser("user01", "password_for_user01")

// protect your existing http.HandlerFunc
proctedHandler := auth.Auth(func(your_http_handler_func))

```

## How to load passwords from a file?
The lib also supports loading multiple users from a file.
```
auth := basicauth.NewBasicAuth()
auth.LoadUsersFromFile("your_pass_file")
```
The password file format is:

```
user01:base64(sha256(password_for_user01))
user02:base64(sha256(password_for_user02))
...
```

## How to generate a password file?
You can use any tools to generate a password file as long as the result file
follows the format mentioned above.

### shell script
```
hashedPwd=$(echo -n $your_password | sha256sum | cut -f1 -d" " | xxd -r -p | base64)
echo $username:$your_password >> your_pass_file
```

### built-in tool
This project has a built-in tool as well in the "cmd" folder.
```
go run ./cmd/adduser.go usernmae password >> your_password_file 
```

## Example code
```go
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
	http.HandleFunc("/protected", auth.Auth(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this page is protected by basic auth.")
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
