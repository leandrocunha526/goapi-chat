package register

import (
	"context"
	"database/sql"

	"github.com/leandrocunha526/goapi-chat/model/api"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, db *sql.DB) error {
	ctx := context.Background()
	user := new(api.RegisterRequest)

	if err := c.BodyParser(user); err != nil {
		panic(err)
	}
	res := Service(db, ctx, *user)
	return c.JSON(res)
}
