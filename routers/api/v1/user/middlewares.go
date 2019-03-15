package user

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	identityKey = "id"
	// Auth 认证中间件
	Auth *jwt.GinJWTMiddleware
	err  error
)

func init() {
	// the jwt middleware
	fmt.Println("init jwt")
	Auth, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// 登录时调用，可将载荷添加到token中
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("调用：PayloadFunc")
			if v, ok := data.(*user); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{
				identityKey: data,
			}
		},
		// 验证登录状态
		IdentityHandler: func(c *gin.Context) interface{} {
			fmt.Println("调用：IdentityHandler")
			claims := jwt.ExtractClaims(c)
			// return &User{
			// 	Username: claims["id"].(string),
			// }
			fmt.Println(claims[identityKey])
			return claims[identityKey]
		},
		// 验证登录
		Authenticator: func(c *gin.Context) (interface{}, error) {
			login := &login{}
			if err := c.ShouldBind(login); err != nil {
				return "", err
			}
			user, msg, result := login.validator()

			if result {
				session := sessions.Default(c)
				session.Set("role", user.Role)
				session.Save()
				return &user, nil
			}

			return nil, errors.New(msg)
		},
		// 鉴权成功后执行
		Authorizator: func(data interface{}, c *gin.Context) bool {
			session := sessions.Default(c)
			session.Set("userInfo", data)
			session.Save()
			return true
		},
		// 登录成功的回调函数
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "login success!",
			})
		},
		// 登录失效时的回调函数
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,

		// Optionally return the token as a cookie
		SendCookie: true,
	})
}
