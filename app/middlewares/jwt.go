package middlewares

import (
	"belajar/bareng/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: 	[]byte(config.SECRET_JWT),
		SigningMethod: 	"HS256",
	})
}

func CreateToken(UserId int) (string, error){
	claims := jwt.MapClaims{}
	claims["authorized"]=true
	claims["userId"]=UserId
	claims["exp"]=time.Now().Add(time.Hour*48).Unix()
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractTokenUserId(e echo.Context) uint{
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}
	return 0
}