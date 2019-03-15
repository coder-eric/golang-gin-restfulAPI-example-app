package main

import (
	"fmt"
	_ "golang-gin-restfulAPI-example-app/common/validator"
	"golang-gin-restfulAPI-example-app/conf"
	"golang-gin-restfulAPI-example-app/routers"
)

// @title yottachain API
// @version 1.0
// @description yottachain 的API接口.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	port, err := conf.Cfg.Section("server").GetKey("HTTP_PORT")
	if err != nil {
		fmt.Println("loade config fail")
	}

	r := routers.InitRouter()
	r.Run(":" + port.String())
}
