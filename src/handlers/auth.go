package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"access_token": "123",
		"user": map[string]interface{}{
			"id":        1,
			"email":     "user@example.com",
			"full_name": "User",
		},
	})
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
