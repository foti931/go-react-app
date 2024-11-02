package usecase

import (
	"errors"
	"go-rest-api/models"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserUsecase interface {
	SignUp(user *models.User) (models.UserResponse, error)
	Login(user *models.User) (string, error)
	UpdateUser(user *models.User) error
	PasswordResetRequest(email string) (string, error)
}

type UserUseCase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
	pr repository.IPasswordRepository
}

// UpdateUser PasswordReset implements IUserUseCase.
func (u *UserUseCase) UpdateUser(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return u.ur.UpdateUser(user)
}

// PasswordResetRequest implements IUserUsecase.
func (u *UserUseCase) PasswordResetRequest(email string) (string, error) {

	//ユーザーが存在するか確認
	existsUser := &models.User{}
	if err := u.ur.GetUserByEmail(existsUser, email); err != nil {
		slog.Error(err.Error())
		return "", err
	}

	//ユーザーが存在しない場合
	if existsUser == nil {
		slog.Error("user not exists")
		return "", errors.New("user not exists")
	}

	//パスワードリセット用のjwtトークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		// 有効期限を5分に設定
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	//トークンを文字列に変換
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	//パスワードリセットリクエストを作成
	if err := u.pr.CreatePasswordResetRequest(existsUser, tokenString); err != nil {
		slog.Error(err.Error())
		return "", err
	}

	return tokenString, nil
}

func (u *UserUseCase) SignUp(input *models.User) (models.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return models.UserResponse{}, err
	}

	newUser := models.User{
		Email:    input.Email,
		Password: string(hash),
	}

	//ユーザー情報の取得
	existsUser := models.User{}
	if err := u.ur.GetUserByEmail(&existsUser, input.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.UserResponse{}, err
	}

	//ユーザーが存在する場合
	if existsUser != (models.User{}) {
		return models.UserResponse{}, errors.New("user already exists")
	}

	if err := u.ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}

	resUser := models.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (u *UserUseCase) Login(input *models.User) (string, error) {
	storedUser := models.User{}

	//ユーザー情報の取得
	if err := u.ur.GetUserByEmail(&storedUser, input.Email); err != nil {
		return "", err
	}

	//パスワードの比較
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password)); err != nil {
		return "", err
	}

	//トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		// 有効期限を12時間に設定
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewUserUseCase(ur repository.IUserRepository, pr repository.IPasswordRepository, uv validator.IUserValidator) IUserUsecase {
	return &UserUseCase{ur: ur, pr: pr, uv: uv}
}
