package router

import (
	"fmt"
	"net/http"
	"github.com/gmae199boy/avi_golang/handler"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	var handler handler.UserHandler = handler.UserHandler{}
	switch r.Method {
	case "GET": {
		handler.ReadUser(w, r)
	}; break
	case "POST": {
		// fmt.Println(r.URL.Query())
		// Create User Data Method
		handler.CreateUser(w, r)
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