package main

import (
	"golang_article/config"
	"golang_article/delivery/controller"
	"golang_article/delivery/http"
	"golang_article/repository"
	"golang_article/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDBConnection()
	defer config.CloseDBConnection(db)

	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	articleRepository := repository.NewArticleRepository(db)

	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository)
	articleService := service.NewArticleService(articleRepository)
	authService := service.NewAuthService(userRepository)

	authController := controller.NewAuthController(authService, jwtService)
	userController := controller.NewUserController(userService, jwtService)
	articleController := controller.NewArticleController(articleService, jwtService)

	http.NewArticleHandler(r, authController, userController, articleController, jwtService)

	r.Run()
}
