package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	config2 "github.com/leandrocunha526/goapi-chat/app/config"
	"github.com/leandrocunha526/goapi-chat/model/api"
	"strings"
)

func JWTAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization", "")

	if !strings.Contains(authHeader, "Bearer") || authHeader == "" {
		return c.JSON(api.BaseResponse{
			Success: false,
			Code:    401,
			Message: "Invalid headers authorization",
		})
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
	config, errConfig := config2.LoadConfig()
	if errConfig != nil {
		return c.JSON(api.BaseResponse{
			Success: false,
			Code:    400,
			Message: "Error when read config",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(config.JwtSecretKey), nil
	})
	if err != nil {
		return c.JSON(api.BaseResponse{
			Success: false,
			Code:    400,
			Message: "Invalid token",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.JSON(api.BaseResponse{
			Success: false,
			Code:    400,
			Message: "Token invalid signature",
		})
	}
	c.Locals("jwt", claims)
	return c.Next()
}
