package controller

import (
	"go-rest-api/models"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAll(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type tackController struct {
	tu usecase.ITaskUsecase
}

// Create implements ITaskController.
func (t *tackController) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	task.UserId = uint(userId.(float64))

	taskRes, err := t.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, taskRes)
}

// Delete implements ITaskController.
func (t *tackController) Delete(c echo.Context) error {
	panic("unimplemented")
}

// Get implements ITaskController.
func (t *tackController) Get(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskId, err := strconv.ParseUint(c.Param("task_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	taskRes, err := t.tu.GetTasksById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, taskRes)
}

// GetAll implements ITaskController.
func (t *tackController) GetAll(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	c.Logger().Info("user_id: ", userId)
	c.Logger().Info("claims: ", claims)

	taskRes, err := t.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, taskRes)
}

// Update implements ITaskController.
func (t *tackController) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskId, err := strconv.ParseUint(c.Param("task_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	task := models.Task{}
	if err = c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	taskRes, err := t.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, taskRes)
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &tackController{tu: tu}
}
