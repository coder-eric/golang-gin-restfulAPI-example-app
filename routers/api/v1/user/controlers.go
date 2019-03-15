package user

import (
	"time"

	"golang-gin-restfulAPI-example-app/common/pkg/e"
	. "golang-gin-restfulAPI-example-app/common/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// @Summary 用户注册
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200
// @Failure 500
// @Router /api/v1/tags [post]
func register(c *gin.Context) {
	u := &user{}

	if err := c.ShouldBind(u); err != nil {
		RES(c, e.INVALID_PARAMS, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err = u.getOneByUsername(u.Username); err != nil {
		err = u.add()
		if err != nil {
			RES(c, e.ERROR, gin.H{
				"message": err.Error(),
			})
		} else {
			RES(c, e.SUCCESS, gin.H{})
		}
	} else {
		RES(c, e.ERROR, gin.H{
			"message": "用户名已存在！",
		})
	}
}

// 获取用户列表
func getUserList(c *gin.Context) {
	users := &users{}
	err := users.getAll()
	if err != nil {
		RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		RES(c, e.SUCCESS, gin.H{
			"data": users,
		})
	}
}

// 更新用户信息
func updateUserByID(c *gin.Context) {
	id := c.Param("id")
	sex := c.PostForm("sex")
	user := &user{}
	user.getOneByID(id)
	user.Sex = sex
	user.UpdateTime = time.Now().UnixNano() / int64(time.Millisecond)
	err := user.update()
	if err != nil {
		RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		RES(c, e.SUCCESS, gin.H{})
	}
}

// 删除用户信息
func deleteUserByID(c *gin.Context) {
	user := &user{
		ID: bson.ObjectIdHex(c.Param("id")),
	}
	err := user.delete()
	if err != nil {
		RES(c, e.ERROR, gin.H{
			"message": err.Error(),
		})
	} else {
		RES(c, e.SUCCESS, gin.H{})
	}
}
