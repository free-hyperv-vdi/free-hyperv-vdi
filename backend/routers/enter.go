package routers

import (
	"gin-vue/api/utils"
	"gin-vue/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	})
	cookieCheckMid := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if c.Request.URL.Path == "/api/cloud/v1/login" || c.Request.URL.Path == "/api/cloud/v1/reset_password" {
				c.Next()
			} else {
				
				accessToken, _ := c.Cookie("accessToken")
				userId, _ := c.Cookie("userId")
				checked := utils.CheckAccessToken(accessToken, userId)
				if checked {
					c.Next()
				} else {
					c.AbortWithStatus(http.StatusUnauthorized)
				}

			}

		}
	}
	router.Use(cookieCheckMid())

	router.Use(Cors())
	UsersRouter(router)
	DevicesRouter(router)
	DisksRouter(router)
	GpusRouter(router)
	return router
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
