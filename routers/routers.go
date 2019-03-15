package routers

import (
	"fmt"
	"golang-gin-restfulAPI-example-app/common/middlewares"
	"golang-gin-restfulAPI-example-app/conf"
	v1 "golang-gin-restfulAPI-example-app/routers/api/v1"

	_ "golang-gin-restfulAPI-example-app/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.Default()

	setUpConfig(router)
	setUpRouter(router)

	return router
}

// 初始化应用设置
func setUpConfig(router *gin.Engine) {
	// 设置静态文件处理
	router.Static("/assets", "./web/dist")
	router.StaticFile("/app.html", "./web/dist/index.html")

	// 使用swagger自动生成接口文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 使用session中间件
	sec, err := conf.Cfg.GetSection("database")
	if err != nil {
		fmt.Println("loade config fail")
	}
	host := sec.Key("HOST").String()
	connection, err := mgo.Dial(host)
	if err != nil {
		// handle err
	}
	collection := connection.DB("test").C("sessions")
	store := mongo.NewStore(collection, 3600, true, []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// 使用权限管理中间件
	e := casbin.NewEnforcer("conf/authz/model.conf", "conf/authz/policy.csv")
	router.Use(middlewares.NewAuthorizer(e))
}

// 设置路由
func setUpRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
}
