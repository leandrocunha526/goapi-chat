package login

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/leandrocunha526/goapi-chat/model/api"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	ctx := context.Background()
	user := new(api.LoginRequest)
	if err := c.BodyParser(user); err != nil {
		panic(err)
	}
	res := Service(ctx, db, *user)
	return c.JSON(res)
}
