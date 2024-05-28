package controller

import (
	"go-rest-api/models"
	"go-rest-api/usecase"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func (u *userController) SignUp(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(&user); err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := u.uu.SignUp(user)
	if err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *userController) LogIn(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := u.uu.Login(user)
	if err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusNotFound, err.Error())
	}
	cookie := new(http.Cookie)
	setCookie(cookie, c, tokenString)

	return c.NoContent(http.StatusOK)
}

func (u *userController) LogOut(c echo.Context) error {

	cookie := new(http.Cookie)
	resetCookie(cookie, c)
	return c.NoContent(http.StatusOK)
}

func (u *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu: uu}
}

func setCookie(cookie *http.Cookie, c echo.Context, tokenString string) {
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}

func resetCookie(cookie *http.Cookie, c echo.Context) {
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
}
