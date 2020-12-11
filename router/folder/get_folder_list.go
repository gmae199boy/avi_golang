package folder

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gmae199boy/avi_golang/model"
	"github.com/gmae199boy/avi_golang/db"
	"github.com/gmae199boy/avi_golang/fileSystem"
)

func GetFolderList (w http.ResponseWriter, r *http.Request) {
	// 유저 정보를 body에서 꺼내와 json 파싱한다.
	var user model.User
	e := json.NewDecoder(r.Body).Decode(&user)
	if e != nil {
		fmt.Println(e)
		fmt.Fprintf(w, e.Error())
		return
	}

	// (users collection) User DB 에서 유저 정보가 있는지 찾은 후, 없다면 유저 정보 저장
	var userDB = new(db.UserDB)
	userDB.Connect()
	u, dbE := userDB.FindOneUser(user.Name)
	if dbE != nil {
		fmt.Println(dbE.Error())
		fmt.Fprintf(w, dbE.Error())
		return
	} else if u == nil && dbE == nil {
		fmt.Println("유저 정보가 없어서 폴더를 만들 수 없음")
		fmt.Println("유저 정보가 없어서 폴더를 만들 수 없음")
		return
	}

	fileSystem.OpenFolder("./assets/kim")
}