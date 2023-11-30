package routes

import (
	_ "github.com/ariwiraa/my-gram/docs"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/middlewares"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Mygram
// @version 1.0
// @description This is a Final Project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description	How to input in swagger : 'Bearer <insert_your_token_here>'
func NewRouter(userHandler handler.UserHandler, photoHandler handler.PhotoHandler, commentHandler handler.CommentHandler, likesHandler handler.UserLikesPhotosHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.PostUserRegisterHandler)
	router.POST("/signin", userHandler.PostUserLoginHandler)
	router.PUT("/refresh", userHandler.PutAccessTokenHandler)
	router.DELETE("/signout", middlewares.Authentication(), userHandler.LogoutHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	photo := router.Group("/photos")
	{
		// Photo
		photo.Use(middlewares.Authentication())
		photo.POST("", photoHandler.PostPhotoHandler)
		photo.GET("/all", photoHandler.GetPhotosHandler)
		photo.GET("", photoHandler.GetPhotosByUserIdHandler)
		photo.GET("/:id", photoHandler.GetPhotoHandler)
		photo.PUT("/:id", photoHandler.PutPhotoHandler)
		photo.DELETE("/:id", photoHandler.DeletePhotoHandler)

		// Likes Photo
		photo.POST("/:id/likes", likesHandler.PostLikesHandler)
		photo.GET("/:id/likes", likesHandler.GetUsersWhoLikedPhotosHandler)

		// Comments
		photo.POST("/:id/comments", commentHandler.PostCommentHandler)
		photo.GET("/:id/comments", commentHandler.GetCommentsHandler)
		photo.GET("/:id/comments/:commentId", commentHandler.GetCommentHandler)
		photo.PUT("/:id/comments/:commentId", commentHandler.PutCommentHandler)
		photo.DELETE("/:id/comments/:commentId", commentHandler.DeleteCommentHandler)
	}

	me := router.Group("/me")
	{
		me.Use(middlewares.Authentication())
		me.GET("/liked/photos", likesHandler.GetPhotosLikedHandler)
	}

	return router
}
