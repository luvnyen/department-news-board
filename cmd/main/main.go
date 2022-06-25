package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luvnyen/department-news-board/cmd/config"

	"github.com/luvnyen/department-news-board/middleware"
	controller "github.com/luvnyen/department-news-board/pkg/controllers"
	"github.com/luvnyen/department-news-board/repository"
	"github.com/luvnyen/department-news-board/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	newsRepository repository.NewsRepository = repository.NewNewsRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	newsService    service.NewsService       = service.NewNewsService(newsRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	newsController controller.NewsController = controller.NewNewsController(newsService, jwtService)
)

func main() {
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", authController.Logout)
		authRoutes.POST("/register", authController.Register)
	}

	newsRoutes := router.Group("api/news")
	{
		newsRoutes.GET("/", newsController.All)
		newsRoutes.GET("/:id", newsController.FindByID)
		newsRoutes.POST("/", newsController.InsertNews, middleware.AuthorizeJWT(jwtService))
		newsRoutes.PUT("/:id", newsController.UpdateNews, middleware.AuthorizeJWT(jwtService))
		newsRoutes.DELETE("/:id", newsController.DeleteNews, middleware.AuthorizeJWT(jwtService))
		newsRoutes.GET("/download/:id", newsController.DownloadNews, middleware.AuthorizeJWT(jwtService))
	}

	router.Run("127.0.0.1:3000")
}
