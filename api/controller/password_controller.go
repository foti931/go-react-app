package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
	"go-rest-api/models"
	"go-rest-api/usecase"
	"log/slog"
	"net/http"
	"os"
)

type IPasswordController interface {
	ResetPasswordRequest(c echo.Context) error
	ResetPassword(c echo.Context) error
}

type PasswordController struct {
	pu  usecase.IUserUsecase
	mu  usecase.IMailUsecase
	pru usecase.IPasswordResetUseCase
}

// ResetPassword implements IPasswordController.

func (p *PasswordController) ResetPassword(c echo.Context) error {
	tokenString := c.QueryParam("token")
	password := c.FormValue("password")

	// クエリパラメータなし
	if tokenString == "" {
		return c.JSON(http.StatusBadRequest, "必要な情報が不足しています。再度パスワードリセットを実施してください。")
	}

	// トークン検証およびパース
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムの検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		// 有効期限切れ
		if errors.Is(err, jwt.ErrTokenExpired) {
			return c.JSON(http.StatusBadRequest, "リセットの有効期限が切れています。再度パスワードリセットを実施してください。")
		}

		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return c.JSON(http.StatusBadRequest, "トークン認証エラー")
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusBadRequest, "再度パスワードリセットを実施してください。")
	}

	request := models.PasswordReset{
		Email: tokenClaims["email"].(string),
		Token: tokenString,
	}

	user, err := p.pru.GetPasswordResetRequest(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Password = password
	// パスワード更新
	if err := p.pu.UpdateUser(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// ResetPasswordRequest implements IPasswordController.
func (p *PasswordController) ResetPasswordRequest(c echo.Context) error {

	slog.Info("ResetPasswordRequest start")
	passReset := models.PasswordReset{}
	if err := c.Bind(&passReset); err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	slog.Info("ResetPasswordRequest start updatedb")
	token, err := p.pu.PasswordResetRequest(passReset.Email)
	if err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	env, ok := os.LookupEnv("FE_URL")
	if !ok {
		slog.Error("FE_URL env variable not set")
		return c.JSON(http.StatusInternalServerError, "FE_URL env variable not set")
	}
	body := fmt.Sprintf(
		`パスワードをリセットします。
以下のURLをクリックしてください。
%v/password/reset?token=%v

有効期限は5分間です。
`, env, token)

	if err := p.mu.SendMail(passReset.Email, "パスワードリセット", token, body); err != nil {
		slog.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	slog.Info("ResetPasswordRequest End")
	return c.NoContent(http.StatusOK)
}

func NewPasswordController(pu usecase.IUserUsecase, mu usecase.IMailUsecase, pru usecase.IPasswordResetUseCase) IPasswordController {
	return &PasswordController{pu: pu, mu: mu, pru: pru}
}
