package decoder

import (
	"net/http"
	"encoding/json"
	"errors"
	"github.com/gmae199boy/avi_golang/model"
)

func UserData(u *http.Request) (*model.User, error) {
	var data model.User

	e := json.NewDecoder(u.Body).Decode(&data)
	if e != nil {
		return nil, errors.New("Error User data parsing: inspectAndDecodeUserData")
	}

	// if data.Name == nil || data.Password == nil || data.Email == nil {
	// 	fmt.Println("User Data has empty")
	// 	return nil, false
	// }
	return &data, nil
}