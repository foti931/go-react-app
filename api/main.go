package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB()

	passwordResetRepository := repository.NewPasswordRepository(db)
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUseCase(userRepository, passwordResetRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	// passwordResetValidator := validator.NewPasswordResetValidator()
	mailRepository := repository.NewMailRepository()
	mailUseCase := usecase.NewMailUsecase(mailRepository)
	passwordRepository := repository.NewPasswordRepository(db)
	passwordResetUseCase := usecase.NewPasswordResetUseCase(passwordRepository)
	passwordController := controller.NewPasswordController(userUsecase, mailUseCase, passwordResetUseCase)

	server := router.NewRouter(userController, taskController, passwordController)

	server.Logger.Fatal(server.Start(":8000"))
}
