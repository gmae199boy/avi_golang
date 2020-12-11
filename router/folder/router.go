package folder

import (
	"fmt"
	"net/http"
)

func FolderRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": {
		GetFolderList(w, r)
	}; break
	case "POST": {
		CreateFolder(w, r)
	}; break
	case "PUT": {

	}; break
	case "DELETE": {
		DeleteFolder(w, r)
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