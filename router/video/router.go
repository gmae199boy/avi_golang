package video

import (
	// "fmt"
	"net/http"
	// "github.com/davecgh/go-spew/spew"
)

func VideoRouter(w http.ResponseWriter, r *http.Request) {
	// spew.Dump(r)
	switch r.Method {
	case "GET": {
		
	}; break
	case "POST": {
		UploadVideo(w, r)
	}; break
	}
}