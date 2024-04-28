package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"go-rest-api/models"
	"go-rest-api/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type IUserUsecase interface {
	SignUp(user models.User) (models.UserResponse, error)
	Login(user models.User) (string, error)
}

type UserUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &UserUsecase{ur: ur}
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
