package usecase

import (
	"errors"
	"go-rest-api/models"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserUsecase interface {
	SignUp(user models.User) (models.UserResponse, error)
	Login(user models.User) (string, error)
}

type UserUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &UserUsecase{ur: ur, uv: uv}
}

func (u UserUsecase) SignUp(input models.User) (models.UserResponse, error) {
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

func (u UserUsecase) Login(input models.User) (string, error) {
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
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
