package route

import (
	"net/http"

	"github.com/bactruongvan17/taskhub-userservice/src/dbs"
	"github.com/bactruongvan17/taskhub-userservice/src/handlers"
	"github.com/gin-gonic/gin"
)

func NewService() *gin.Engine {
	// load db
	db := dbs.NewPostgresClient()

	r := gin.Default()

	// handler
	migrationHandler := handlers.NewMigrationHandler(db)
	authHandler := handlers.NewAuthHandler(db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/internal/migrate", migrationHandler.Migrate)

	// authHandler
	r.POST("/auth/signin", authHandler.SignIn)
	r.POST("/auth/signup", authHandler.SignUp)

	return r
}
