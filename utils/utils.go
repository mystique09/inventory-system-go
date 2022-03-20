package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * 365 * time.Hour) // 1 year
	cookie.Path = "/"
	cookie.HttpOnly = true
	//cookie.Secure = true

	c.SetCookie(cookie)
}

func CreateJwt(payload UserJwtPayload) (string, error) {
	claims := JwtClaims{
		payload,
		jwt.StandardClaims{
			Id:        payload.Id.String(),
			ExpiresAt: time.Now().Add(24 * time.Minute).Unix(),
		},
	}

	raw_token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := raw_token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}

type (
	UserJwtPayload struct {
		Id       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
	}

	JwtClaims struct {
		Payload UserJwtPayload
		jwt.StandardClaims
	}
)
