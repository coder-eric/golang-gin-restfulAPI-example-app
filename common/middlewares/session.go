package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

// SessionsMiddleware .
func SessionsMiddleware(c *gin.Context) gin.HandlerFunc {
	session, err := mgo.Dial("localhost:27017/test")
	if err != nil {
		// handle err
	}

	collection := session.DB("test").C("sessions")
	store := mongo.NewStore(collection, 3600, true, []byte("secret"))
	return sessions.Sessions("mysession", store)
}
