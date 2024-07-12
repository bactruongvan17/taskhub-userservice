package route

import (
	"net/http"

	"github.com/bactruongvan17/taskhub-userservice/src/pkg/dbs"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/handlers"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/repo"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/service"
	"github.com/gin-gonic/gin"
)

func NewService() *gin.Engine {
	// load db
	db := dbs.NewPostgresClient()

	r := gin.Default()

	// repo
	repoPG := repo.NewPGRepo(db)

	// service
	authService := service.NewAuthService(repoPG)

	// handler
	migrationHandler := handlers.NewMigrationHandler(db)
	authHandler := handlers.NewAuthHandler(authService)

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
