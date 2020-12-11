package db

import (
	"fmt"
	"errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gmae199boy/avi_golang/model"
	"github.com/davecgh/go-spew/spew"
)

type UserDB struct {
	coll *mongo.Collection
}

func (u *UserDB) Connect() {
	db := GetDB()
	u.coll = db.Collection("users")
}

func (u *UserDB) CreateUser(user *model.User) error {
	userdata, fe := u.FindOneUser(user.Name)
	if userdata != nil || fe != nil {
		return errors.New("DB Error: CreateUser func -> User data founded or DB Error")
	}

	aa, ie := u.coll.InsertOne(context.TODO(), user)
	if ie != nil {
		return ie
	}
	spew.Dump(aa)
	fmt.Println("success save")
	return nil
}


// findOneUser: 유저 하나를 찾는다
/*
	return case 1: return nil, nil : 조건에 맞는 유저 정보가 없을 때
	return case 2: return nil, e   : 쿼리에서 에러가 나올 때
	return case 3: return data, nil: 데이터를 찾았을 때
*/
func (u *UserDB) FindOneUser(name string) (*model.User, error){
	var t model.User
	e := u.coll.FindOne(context.TODO(), bson.M{"name": name}).Decode(&t)
	if e == mongo.ErrNoDocuments {return nil, nil;}
	if e != nil {return nil, e;}
	return &t, nil
}