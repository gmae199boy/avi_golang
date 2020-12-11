package user

import (
	"fmt"
	"net/http"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": {
		// ReadUser(w, r)
	}; break
	case "POST": {
		// fmt.Println(r.URL.Query())
		// Create User Data Method
		CreateUser(w, r)
	}; break
	case "PUT": {

	}; break
	case "DELETE": {

	}; break
	default: {
		fmt.Println("Method Error!! not GET,POST,PUT,DELETE")
	}; break
	}
}