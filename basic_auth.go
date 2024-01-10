package basicauth

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

type BasicAuth struct {
	//map[username]password
	Userdb map[string]string
}

func NewBasicAuth() *BasicAuth {
	return &BasicAuth{
		Userdb: make(map[string]string),
	}
}

func (auth *BasicAuth) EncryptPass(pass string) string {
	sum := sha256.Sum256([]byte(pass))
	ecoded := base64.StdEncoding.EncodeToString(sum[:])
	return ecoded
}

func (auth *BasicAuth) VerifyUser(username, password string) bool {
	return auth.Userdb[username] == auth.EncryptPass(password)
}

func (auth *BasicAuth) AddUser(username, password string) {
	if auth.Userdb == nil {
		auth.Userdb = make(map[string]string)
	}
	hashedPwd := auth.EncryptPass(password)
	auth.Userdb[username] = hashedPwd
}

/*
password file format
username1:base64(sha256(pass1))
username2:base64(sha256(pass2))
*/
func (auth *BasicAuth) LoadUsersFromFile(fname string) error {
	if auth.Userdb == nil {
		auth.Userdb = make(map[string]string)
	}
	content, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	strContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(strContent, "\n")
	for _, line := range lines {
		pair := strings.Split(line, ":")
		if len(pair) < 2 {
			continue
		}
		username := pair[0]
		hashedPwd := pair[1]
		auth.Userdb[username] = hashedPwd
	}
	return nil
}

func (auth *BasicAuth) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok && auth.VerifyUser(username, password) {
			next(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate", `Basic`)
		http.Error(w, "not allowed", 401)
	}
}
