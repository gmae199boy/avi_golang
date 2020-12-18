package router

import (
	// "fmt"
	"net/http"
	// "github.com/davecgh/go-spew/spew"
	"github.com/gmae199boy/avi_golang/handler"
)

func VideoRouter(w http.ResponseWriter, r *http.Request) {
	var handler handler.VideoHandler = handler.VideoHandler{}
	// spew.Dump(r)
	switch r.Method {
	case "GET": {
		
	}; break
	case "POST": {
		handler.UploadVideo(w, r)
	}; break
	}
}