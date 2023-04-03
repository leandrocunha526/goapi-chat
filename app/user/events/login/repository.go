package login

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/leandrocunha526/goapi-chat/app/helper"
	"github.com/leandrocunha526/goapi-chat/model/entity"
)

func Repository(ctx context.Context, tx *sql.Tx, username string) (*entity.User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, query, username)
	helper.PanicError(err)
	var user entity.User
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Image)
		if err != nil {
			fmt.Println(err)
			return new(entity.User), err
		}
		log.Print(user)
	}
	return &user, nil
}
