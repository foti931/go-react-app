package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	// /tasks
	setTaskEndpoint(e, tc)

	return e
}

func setTaskEndpoint(e *echo.Echo, tc controller.ITaskController) {
	taskGroup := e.Group("/tasks")
	taskGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	// GET
	taskGroup.GET("", tc.GetAll)
	taskGroup.GET("/:task_id", tc.Get)
	// POST
	taskGroup.POST("", tc.Create)
	// PUT
	taskGroup.PUT("/:task_id", tc.Update)
	// DELETE
	taskGroup.DELETE("/:task_id", tc.Delete)
}
