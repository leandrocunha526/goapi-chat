package register

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/leandrocunha526/goapi-chat/app/helper"
	"github.com/leandrocunha526/goapi-chat/model/api"
	"github.com/leandrocunha526/goapi-chat/model/entity"
	"golang.org/x/crypto/bcrypt"
)

func Service(db *sql.DB, ctx context.Context, request api.RegisterRequest) *api.RegisterResponse {
	tx, err := db.Begin()
	helper.PanicError(err)
	defer helper.RollbackErr(tx)
	bytes, errHash := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicError(errHash)
	user := entity.User{
		Username: request.Username,
		Password: string(bytes),
		Email:    request.Email,
		Image:    request.Image,
	}
	result, errQuery := Repository(ctx, tx, user)
	var baseResponse api.BaseResponse
	if errQuery != nil {
		fmt.Println(errQuery)
		if strings.Contains(errQuery.Error(), "duplicate") {
			baseResponse = api.BaseResponse{
				Success: false,
				Code:    400,
				Message: "User was registered",
			}
			return &api.RegisterResponse{
				BaseResponse: &baseResponse,
			}
		}
		baseResponse = api.BaseResponse{
			Success: false,
			Code:    500,
			Message: "Something wrong",
		}
		return &api.RegisterResponse{
			BaseResponse: &baseResponse,
		}
	}
	baseResponse = api.BaseResponse{
		Success: true,
		Code:    201,
		Message: "Register is success",
	}
	return &api.RegisterResponse{
		BaseResponse: &baseResponse,
		Data: &api.RegisterResponseData{
			Id:       result.Id,
			Username: result.Username,
			Email:    request.Email,
			Image:    result.Image,
		},
	}
}
