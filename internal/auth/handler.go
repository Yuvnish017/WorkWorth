package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var request SignUpRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	response, err := h.authService.SignUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	response, err := h.authService.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
