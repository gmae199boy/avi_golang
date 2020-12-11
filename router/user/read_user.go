package user

import (
	"fmt"
	// "errors"
	"net/http"
	"encoding/json"
	// "github.com/gmae199boy/avi_golang/db"
	"github.com/gmae199boy/avi_golang/model"
	"github.com/davecgh/go-spew/spew"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	e := json.NewDecoder(r.Body).Decode(&user)
	if e != nil {
		fmt.Println(e)
		fmt.Fprintf(w, e.Error())
		return
	}

	spew.Dump(user)
}