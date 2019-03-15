package user

import (
	"errors"
	"golang-gin-restfulAPI-example-app/common/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var userCollection *mgo.Collection

func init() {
	dbConnect := db.NewConnection()
	userCollection = dbConnect.Use("test", "user")
}

// User 是用户类型
type user struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string        `json:"username" form:"username" validate:"required,email" bson:"username"`
	Password   string        `json:"password" form:"password" validate:"required" bson:"password"`
	Sex        string        `json:"sex" form:"sex" bson:"sex"`
	UpdateTime int64         `json:"updateTime" bson:"updateTime"`
	Role       string        `json:"role" form:"role" bson:"role"`
}

// Users 是用户列表
type users []user

// Login is login struct
type login struct {
	Username string `form:"username" validate:"required" bson:"username"`
	Password string `form:"password" validate:"required" bson:"password"`
}

func checkMongoDBNotNull() error {
	if userCollection == nil {
		return errors.New("mongoDB client is null please set it before use")
	}
	return nil
}

// Add 添加用户
func (u *user) add() error {
	err := checkMongoDBNotNull()
	err = userCollection.Insert(u)
	return err
}

// Validator .
func (login *login) validator() (*user, string, bool) {
	user := &user{}
	err := userCollection.Find(bson.M{"username": login.Username}).One(user)
	var msg string
	if err != nil {
		msg = "没有该账户！"
		return nil, msg, false
	}

	if user.Password != login.Password {
		msg = "密码错误！"
		return nil, msg, false
	}

	msg = "登录成功！"
	return user, msg, true
}

// GetOneByUsername 根据username查询
func (u *user) getOneByUsername(username string) error {
	query := bson.M{"username": username}
	err := userCollection.Find(query).One(u)
	return err
}

// GetOneByID 根据Id查询
func (u *user) getOneByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := userCollection.Find(query).One(u)
	return err
}

// Update 更新用户信息
func (u *user) update() error {
	query := bson.M{"_id": u.ID}
	err := userCollection.Update(query, u)
	return err
}

// GetAll 获取用户列表
func (us *users) getAll() error {
	err := checkMongoDBNotNull()
	err = userCollection.Find(nil).All(us)
	return err
}

// Delete 通过Id删除用户
func (u *user) delete() error {
	err := userCollection.RemoveId(u.ID)
	return err
}
