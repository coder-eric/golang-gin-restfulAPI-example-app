package db

import (
	"fmt"
	"golang-gin-restfulAPI-example-app/conf"

	"gopkg.in/mgo.v2"
)

// Connection 数据库
type Connection struct {
	session *mgo.Session
}

// NewConnection 新建数据库连接
func NewConnection() (conn *Connection) {
	sec, err := conf.Cfg.GetSection("database")
	if err != nil {
		fmt.Println("loade config fail")
	}
	host := sec.Key("HOST").String()
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("connect mongoDB fail!")
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	conn = &Connection{session}

	return
}

// Use 切换数据库
func (conn *Connection) Use(dbName, tableName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(tableName)
}

// Close 关闭数据库连接
func (conn *Connection) Close() {
	conn.session.Close()
}
