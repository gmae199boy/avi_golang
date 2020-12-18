package folder

import (
	// "fmt"
	"net/http"
	"time"
	"os"
	"encoding/json"
	"strings"
	"bytes"
	"github.com/gmae199boy/avi_golang/model"
	// "github.com/gmae199boy/avi_golang/db"
	"github.com/gmae199boy/avi_golang/fileSystem"
	"github.com/dgrijalva/jwt-go"
)

func GetFolderList (w http.ResponseWriter, r *http.Request) {
	// // 유저 정보를 body에서 꺼내와 json 파싱한다.
	// var user model.User
	// e := json.NewDecoder(r.Body).Decode(&user)
	// if e != nil {
	// 	fmt.Println(e)
	// 	fmt.Fprintf(w, e.Error())
	// 	return
	// }

	// // (users collection) User DB 에서 유저 정보가 있는지 찾은 후, 없다면 유저 정보 저장
	// var userDB = new(db.UserDB)
	// userDB.Connect()
	// u, dbE := userDB.FindOneUser(user.Name)
	// if dbE != nil {
	// 	fmt.Println(dbE.Error())
	// 	fmt.Fprintf(w, dbE.Error())
	// 	return
	// } else if u == nil && dbE == nil {
	// 	fmt.Println("유저 정보가 없어서 폴더를 만들 수 없음")
	// 	fmt.Println("유저 정보가 없어서 폴더를 만들 수 없음")
	// 	return
	// }

	// token, err := CreateToken(123)
	// if err != nil {
	// 	w.Write([]byte("asd"))
	// }

	// fmt.Println(token)
	var err error
	var folderPath bytes.Buffer
	userName := strings.Join(r.URL.Query()["username"], "")
	folderName := strings.Join(r.URL.Query()["folder"], "")

	folderPath.WriteString(userName)
	folderPath.WriteString("/")
	folderPath.WriteString(folderName)

	fileName, err := fileSystem.ReadNameOfFileList(folderPath.String())
	if err != nil {

	}
	var fileList = make([]model.File, len(fileName))
	for i := 0; i < len(fileName); i++ {
		fileList[i] = model.File{Name: fileName[i]}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&fileList)
}

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	// os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
  }