package v1

import (
	miningmachine "golang-gin-restfulAPI-example-app/routers/api/v1/mining-machine"
	"golang-gin-restfulAPI-example-app/routers/api/v1/user"

	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		// 用户路由
		user.RegisterRouter(v1.Group("/user"))
		// 矿机管理路由
		miningmachine.RegisterRouter(v1.Group("/miningMachine"))
	}
}
