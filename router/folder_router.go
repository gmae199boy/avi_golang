package router

import (
	"fmt"
	"net/http"
	"github.com/gmae199boy/avi_golang/handler"
)

func FolderRouter(w http.ResponseWriter, r *http.Request) {
	var handler handler.FolderHandler = handler.FolderHandler{}
	switch r.Method {
	case "GET": {
		handler.GetFolderList(w, r)
	}; break
	case "POST": {
		handler.CreateFolder(w, r)
	}; break
	case "PUT": {

	}; break
	case "DELETE": {
		handler.DeleteFolder(w, r)
	}; break
	default: {
		fmt.Println("Method Error!! not GET,POST,PUT,DELETE")
	}; break
	}
}


// func FolderRouter(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET": {
		
// 	}; break
// 	case "POST": {
// 		fmt.Println(r.URL.Query())
// 	}; break
// 	case "PUT": {

// 	}; break
// 	case "DELETE": {

// 	}; break
// 	default: {
// 		fmt.Println("Method Error!! not GET,POST,PUT,DELETE")
// 	}; break
// 	}
// }