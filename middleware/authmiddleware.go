package middleware

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	beego.MiddleWare
}

func ValidateToken(ctx *context.Context) {
	if ctx.Input.URL() == "/auth/login" {
		return
	}
	tokenString := ctx.Input.Header("Authorization")
	if tokenString == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Authorization token is required",
		}, true, false)
		return
	}

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	secretKey := beego.AppConfig.String("jwt_secret_key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "Invalid or expired token",
		}, true, false)
		return
	}
}
