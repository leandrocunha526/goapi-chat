package login

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	config2 "github.com/leandrocunha526/goapi-chat/app/config"
	"github.com/leandrocunha526/goapi-chat/app/helper"
	"github.com/leandrocunha526/goapi-chat/model/api"
	"golang.org/x/crypto/bcrypt"
)

type MyJwtClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
}

var (
	AppName               = "Nextchat"
	JwtSigningMethod      = jwt.SigningMethodHS256
	JwtAccessTokenExpired = time.Duration(30) * (time.Hour * 24)
)

func Service(ctx context.Context, db *sql.DB, request api.LoginRequest) *api.LoginResponse {

	tx, err := db.Begin()
	helper.PanicError(err)
	defer helper.RollbackErr(tx)

	var baseResponse api.BaseResponse
	result, errQuery := Repository(ctx, tx, request.Username)

	if errQuery != nil {

		if strings.Contains(errQuery.Error(), "found") {

			baseResponse = api.BaseResponse{
				Success: false,
				Code:    401,
				Message: errQuery.Error(),
			}

			return &api.LoginResponse{
				BaseResponse: &baseResponse,
			}
		}

		baseResponse = api.BaseResponse{
			Success: false,
			Code:    401,
			Message: "error when query to database",
		}

		return &api.LoginResponse{
			BaseResponse: &baseResponse,
		}

	}

	fmt.Println(result.Password)
	errComparePass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(request.Password))
	if errComparePass != nil {
		fmt.Println(result.Password)
		return &api.LoginResponse{
			BaseResponse: &api.BaseResponse{
				Success: false,
				Code:    401,
				Message: "Password are invalid",
			},
		}
	}

	baseResponse = api.BaseResponse{
		Success: true,
		Code:    200,
		Message: "Login is success",
	}

	accessTokenClaims := MyJwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    AppName,
			ExpiresAt: time.Now().Add(JwtAccessTokenExpired).Unix(),
		},
		Username: result.Username,
		Email:    result.Email,
		Id:       result.Id,
	}

	refreshTokenClaims := MyJwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    AppName,
			ExpiresAt: time.Now().Add(JwtAccessTokenExpired).Unix(),
		},
		Username: result.Username,
		Email:    result.Email,
		Id:       result.Id,
	}

	accessToken := jwt.NewWithClaims(JwtSigningMethod, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(JwtSigningMethod, refreshTokenClaims)

	config, errConfig := config2.LoadConfig()
	if errConfig != nil {
		return &api.LoginResponse{
			BaseResponse: &api.BaseResponse{
				Success: false,
				Code:    401,
				Message: err.Error(),
			},
		}
	}

	signedAccessToken, errSignedToken := accessToken.SignedString([]byte(config.JwtSecretKey))
	signedRefreshToken, errSignedToken := refreshToken.SignedString([]byte(config.JwtSecretKey))

	if errSignedToken != nil {
		fmt.Println(errSignedToken)
		panic(errSignedToken)
	}

	return &api.LoginResponse{
		BaseResponse: &baseResponse,
		Data: &api.LoginResponseData{
			AccessToken:  signedAccessToken,
			RefreshToken: signedRefreshToken,
		},
	}
}
