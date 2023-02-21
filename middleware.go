package main

import "net/http"

const USERNAME = "dee"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte(`something wrong`))
		return false
	}
	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		w.Write([]byte(`Wrong Credentials`))
		return false
	}
	return true
}

func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`Only GET Allowed`))
		return false
	}
	return true
}
