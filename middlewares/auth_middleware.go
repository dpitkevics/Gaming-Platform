package middlewares

import (
	"github.com/appleboy/gin-jwt"
	"github.com/dpitkevics/GamingPlatform/services"
	"github.com/gin-gonic/gin"
	"time"
)

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "Gaming Platform",
		Key:        []byte("asdjkhjfczoiuxyf*&^YA&*AST%dzx"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, password string, context *gin.Context) (string, bool) {
			user := services.GetUserByUsernameAndPassword(username, password)
			if user != nil {
				return username, true
			}

			return username, false
		},
		Authorizator: func(userId string, context *gin.Context) bool {
			return true
		},
		Unauthorized: func(context *gin.Context, code int, message string) {
			context.SecureJSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	return authMiddleware
}
