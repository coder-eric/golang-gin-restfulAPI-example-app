package miningmachine

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.RouterGroup) {
	// 添加矿机
	r.POST("/item", addItem)
	// 矿机列表
	r.GET("/list", getList)
	// 根具矿机id获取矿机详情
	r.GET("/item/:id", getItemByID)
}
