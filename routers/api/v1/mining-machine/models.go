package miningmachine

import (
	"golang-gin-restfulAPI-example-app/common/db"

	"gopkg.in/mgo.v2/bson"
)

var dbConnect *db.Connection

func init() {
	dbConnect = db.NewConnection()
}

// 矿机
type miningMachine struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	NodeID         string        `json:"nodeId" form:"nodeId" validate:"required" bson:"nodeId"`
	Addrs          []string      `json:"addrs" form:"addrs" bson:"addrs"`
	PubAddr        string        `json:"pubAddr" form:"pubAddr" bson:"pubAddr"`
	ExpirationTime string        `json:"expirationTime" bson:"expirationTime"`
}

// 添加矿机
func (m *miningMachine) add() error {
	c := dbConnect.Use("MiningMachine", "list")
	return c.Insert(m)
}

// 根据Id查询
func (m *miningMachine) getOneByID(id string) error {
	c := dbConnect.Use("MiningMachine", "list")
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	return c.Find(query).One(m)
}

// 根据nodeID查询
func (m *miningMachine) getOneByNodeID(nodeID string) error {
	c := dbConnect.Use("MiningMachine", "list")
	query := bson.M{"nodeId": nodeID}
	return c.Find(query).One(m)
}

// 更新矿机信息
func (m *miningMachine) update() error {
	c := dbConnect.Use("MiningMachine", "list")
	query := bson.M{"_id": m.ID}
	return c.Update(query, m)
}

// 矿机列表
type miningMachineList []miningMachine

// 获取矿机列表
func (ml *miningMachineList) getAll() error {
	c := dbConnect.Use("MiningMachine", "list")
	return c.Find(nil).All(ml)
}

// 获取矿机总数
func (ml *miningMachineList) count() int {
	c := dbConnect.Use("MiningMachine", "list")
	count, _ := c.Find(nil).Count()
	return count
}
