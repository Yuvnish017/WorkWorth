package auth

import (
	"WorkWorth/internal/users"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/Yuvnish017/users"
)

type AccessTokenClaim struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *users.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := AccessTokenClaim{
		UserId: user.ID.Hex(),
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func SignUp(c *gin.Context) {
	var request SignUpRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if request.Password != request.ConfirmPassword {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "password and confirm password do not match"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	request.Password = string(encryptedPassword)

	user := users.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err := user.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	accessToken, err := CreateAccessToken(&user, secret, expiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := LoginResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
