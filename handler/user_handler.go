package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gmae199boy/avi_golang/db"
	// "github.com/gmae199boy/avi_golang/fileSystem"
	"github.com/gmae199boy/avi_golang/model"
	"github.com/davecgh/go-spew/spew"
)

type UserHandler struct {

}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
	dbE := userDB.CreateUser(&user)
	if dbE != nil {
		fmt.Println(dbE)
		fmt.Fprintf(w, dbE.Error())
		return
	}

	// 유저가 생성되었기 때문에 유저에 대한 파일을 생성
	// 0777: chmod -R 777 mod
	// fe := fileSystem.CreateFolder(user.Name, 0777)
	// if fe != nil {
	// 	fmt.Println(fe)
	// 	fmt.Fprintf(w, fe.Error())
	// 	return
	// }
	fmt.Fprintf(w, "create user successful")
	// 저장된 유저 정보를 그대로 리턴한다
	// echo data
	// i, _ := json.Marshal(user)

	// spew.Dump(i)
}

func (h UserHandler) ReadUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	e := json.NewDecoder(r.Body).Decode(&user)
	if e != nil {
		fmt.Println(e)
		fmt.Fprintf(w, e.Error())
		return
	}

	spew.Dump(user)
}