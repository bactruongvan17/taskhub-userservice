package handlers

import (
	"net/http"

	"github.com/bactruongvan17/taskhub-userservice/src/pkg/request"
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var req request.SignInRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Data request invalid",
		})
		return
	}

	res, err := h.service.SignIn(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var req request.SignUpRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Data request invalid",
		})
		return
	}

	err := h.service.SingUp(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
