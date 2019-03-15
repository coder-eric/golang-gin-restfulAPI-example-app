# golang-gin-restfulAPI-example-app



> 一个用go语言基于Gin写的restful风格api服务程序的例子.

## 项目特性
- 基于[gin](https://github.com/gin-gonic/gin)
- 使用[MongoDB](https://github.com/mongodb/mongo)数据库
- [gin-jwt](https://github.com/appleboy/gin-jwt)权限验证
- [gin-sessions](https://github.com/gin-contrib/sessions)
- [gin-authz](https://github.com/gin-contrib/authz)从session里取用户的角色进行权限管理
- 使用[gin-swagger](https://github.com/swaggo/gin-swagger)自动生成api文档
- 将gin默认的validator.v8升级到[validator.v9](https://github.com/go-playground/validator)
- 使用[casbin](https://github.com/casbin/casbin)权限管理
- 使用[go-ini](https://github.com/go-ini/ini)读写配置文件

## 项目目录
```
.
├── server.go                       // 入口文件
├── docs                            // swagger生成的api文档
├── web                             // vue写的前端单页页面
├── common
│   ├── db                          // mongoDB相关
│   ├── utils                       // 公用工具函数
│   ├── pkg                         // 公用包
|   |   └── e
|   |       ├── code.go             // http状态码常量
│   |       └── message.go          // 状态码对应的message常量
│   ├── validator
|   |   ├── custom_validate.go      // 自定义验证器
│   |   └── v8_to_v9.go             // 将gin的默认验证器从v8升级到v9
│   └── middlewares     
|       ├── authz.go                // 角色认证
│       └── session.go              // 使用session
├── conf                            // 应用配置相关文件
|   ├── authz                       
|   |   ├── model.conf              // 权限管理方案配置
│   |   └── policy.csv              // 权限分配表
|   ├── app.ini                     // 应用配置文件
│   └── conf.go                     // 初始化配置文件
└── routers
    ├── routers.go                  // 路由初始化
    └── api                         // api文件
        └── v1                      // api版本v1
            ├── mining-machine      // 矿机模块
            |   ├── models.go       // 模型和数据库操作
            |   ├── controlers.go   // 当前模块的控制器
            |   ├── routers.go      // 当前模块的路由
            |   ├── middlewares.go  // 当前模块的中间件
            |   └── validators.go   // 当前模块的验证器
            └── user                // 用户模块
                ├── models.go       // 模型和数据库操作
                ├── controlers.go   // 当前模块的控制器
                ├── routers.go      // 当前模块的路由
                ├── middlewares.go  // 当前模块的中间件
                └── validators.go   // 当前模块的验证器
    
```

## 联系作者
<img src="qrcode.jpg" width="300" align=center alt="微信：wuhe52"/>