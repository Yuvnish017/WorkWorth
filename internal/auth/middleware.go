package auth

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string, secret string) (int64, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			userID, err := ValidateToken(authToken, secret)
			if err != nil {
				c.JSON(http.StatusUnauthorized, ErrorResponse{Message: err.Error()})
				c.Abort()
				return
			}

			c.Set("x-user-id", userID)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
