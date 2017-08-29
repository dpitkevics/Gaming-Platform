package main

import (
	"github.com/dpitkevics/GamingPlatform/database"
	"github.com/dpitkevics/GamingPlatform/http_handlers"
	"github.com/dpitkevics/GamingPlatform/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	database.InitDatabase()
	database.SeedDatabase()
	defer database.CloseDatabase()

	authMiddleware := middlewares.GetAuthMiddleware()

	router := gin.New()

	api := router.Group("/api/v1")
	{
		public := api.Group("")
		{
			public.GET("/", func(context *gin.Context) {
				context.SecureJSON(http.StatusOK, struct {
					Message string `json:"message"`
				}{
					Message: "Hello, World",
				})
			})

			auth := public.Group("/auth")
			{
				auth.POST("/register", http_handlers.CreateUserAction)
				auth.POST("/login", authMiddleware.LoginHandler)
			}
		}

		private := api.Group("")
		private.Use(authMiddleware.MiddlewareFunc())
		{
			users := private.Group("/users")
			{
				users.GET("/me", http_handlers.GetMeAction)
			}
		}
	}

	router.Run()
}
