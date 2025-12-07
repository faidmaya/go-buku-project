package routers

import (
	"go-buku-project/controllers"
	"go-buku-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(useJWT bool) *gin.Engine {
	r := gin.Default()

	// public
	r.GET("/api/categories", controllers.GetCategories)
	r.GET("/api/categories/:id", controllers.GetCategory)
	r.GET("/api/books", controllers.GetBooks)
	r.GET("/api/books/:id", controllers.GetBook)
	r.GET("/api/categories/:id/books", controllers.GetBooksByCategory)

	// auth route
	r.POST("/api/users/login", controllers.Login)

	// protected group
	var authMiddleware gin.HandlerFunc
	if useJWT {
		authMiddleware = middlewares.JWTMiddleware()
	} else {
		authMiddleware = middlewares.BasicAuthMiddleware()
	}

	a := r.Group("/api")
	a.Use(authMiddleware)
	{
		a.POST("/categories", controllers.CreateCategory)
		a.PUT("/categories/:id", controllers.UpdateCategory)
		a.DELETE("/categories/:id", controllers.DeleteCategory)

		a.POST("/books", controllers.CreateBook)
		a.PUT("/books/:id", controllers.UpdateBook)
		a.DELETE("/books/:id", controllers.DeleteBook)
	}

	return r
}
