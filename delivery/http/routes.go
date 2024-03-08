package http

import (
	"golang_article/delivery/controller"
	"golang_article/middleware"
	"golang_article/service"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	authController    controller.AuthController
	userController    controller.UserController
	articleController controller.ArticleController
	jwtService        service.JWTService
}

func NewArticleHandler(routes *gin.Engine, authCon controller.AuthController, userCon controller.UserController, articleCon controller.ArticleController, jwtServ service.JWTService) {
	handler := &articleHandler{
		authController:    authCon,
		userController:    userCon,
		articleController: articleCon,
		jwtService:        jwtServ,
	}

	//handling request for authentication service
	authRoutes := routes.Group("api/auth")
	{
		authRoutes.POST("/login", handler.authController.Login)
		authRoutes.POST("/register", handler.authController.Register)
	}

	//handling request for user service
	userRoutes := routes.Group("api/users", middleware.AuthorizeJWT(handler.jwtService))
	{
		userRoutes.GET("/", handler.userController.AllUser)
		userRoutes.GET("/profile", handler.userController.Profile)
		userRoutes.PUT("/update", handler.userController.Update)
	}

	//handling request for article service
	articleRoutes := routes.Group("api/articles", middleware.AuthorizeJWT(handler.jwtService))
	{
		articleRoutes.GET("/", handler.articleController.All)
		articleRoutes.POST("/create", handler.articleController.Insert)
		articleRoutes.PUT("/update", handler.articleController.Update)
		articleRoutes.GET("/:id", handler.articleController.FindByID)
		articleRoutes.DELETE("/:id", handler.articleController.Delete)
	}

}
