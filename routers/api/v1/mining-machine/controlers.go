package miningmachine

import (
	"fmt"
	"golang-gin-restfulAPI-example-app/common/pkg/e"
	. "golang-gin-restfulAPI-example-app/common/utils"

	"github.com/gin-gonic/gin"
)

// 添加矿机
func addItem(c *gin.Context) {
	m := &miningMachine{}
	m1 := &miningMachine{}
	if err := c.BindJSON(m); err != nil {
		RES(c, e.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := m1.getOneByNodeID(m.NodeID)
	if err == nil {
		m.ID = m1.ID
		err = m.update()
		if err != nil {
			fmt.Println(err)
			RES(c, e.ERROR, gin.H{})
		} else {
			RES(c, e.SUCCESS, gin.H{})
		}
	} else {
		err = m.add()
		if err != nil {
			fmt.Println(err)
			RES(c, e.ERROR, gin.H{})
		} else {
			RES(c, e.SUCCESS, gin.H{})
		}
	}
}

// 获取矿机列表
func getList(c *gin.Context) {
	mList := &miningMachineList{}
	err := mList.getAll()
	if err != nil {
		RES(c, e.ERROR, gin.H{})
	} else {
		RES(c, e.SUCCESS, gin.H{
			"data":  mList,
			"count": mList.count(),
		})
	}
}

// 通过id获取矿机信息
func getItemByID(c *gin.Context) {
	m := &miningMachine{}
	id := c.Param("id")
	err := m.getOneByID(id)
	if err != nil {
		RES(c, e.ERROR, gin.H{})
	} else {
		RES(c, e.SUCCESS, gin.H{
			"data": m,
		})
	}
}
