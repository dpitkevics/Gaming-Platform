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

			sports := public.Group("/sports")
			{
				sports.GET("", http_handlers.GetAllSportsAction)
				sports.GET("/:sportId", http_handlers.GetSportAction)
			}
		}

		private := api.Group("")
		private.Use(authMiddleware.MiddlewareFunc())
		{
			users := private.Group("/users")
			{
				users.GET("/me", http_handlers.GetMeAction)
			}

			groups := private.Group("/groups")
			{
				groups.POST("", http_handlers.CreateGroupAction)
			}

			invitations := private.Group("/invitations")
			{
				invitations.POST("", http_handlers.CreateGroupInvitationAction)
				invitations.GET("/:invitationId", http_handlers.GetGroupInvitationAction)
				invitations.POST("/:invitationId/accept", http_handlers.AcceptGroupInvitationAction)
				invitations.POST("/:invitationId/decline", http_handlers.DeclineGroupInvitationAction)
			}
		}
	}

	router.Run()
}
