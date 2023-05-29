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
func NewRouter(userHandler handler.UserHandler, photoHandler handler.PhotoHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.PostUserRegisterHandler)
	router.POST("/signin", userHandler.PostUserLoginHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	photo := router.Group("/photo")
	{
		photo.Use(middlewares.Authentication())
		photo.POST("", photoHandler.PostPhotoHandler)
		photo.GET("", photoHandler.GetPhotosHandler)
		photo.GET("/:id", photoHandler.GetPhotoHandler)
		photo.PUT("/:id", middlewares.PhotoAuthorization(), photoHandler.PutPhotoHandler)
		photo.DELETE("/:id", middlewares.PhotoAuthorization(), photoHandler.DeletePhotoHandler)
	}

	// socialMedia := router.Group("/socialmedia")
	// {
	// 	socialMedia.Use(middlewares.Authentication())
	// 	socialMedia.POST("", socialMediaHandler.PostSocialMediaHandler)
	// 	socialMedia.GET("", socialMediaHandler.GetSocialMediasHandler)
	// 	socialMedia.GET("/:id", socialMediaHandler.GetSocialMediaHandler)
	// 	socialMedia.PUT("/:id", middlewares.SocialMediaAuthorization(), socialMediaHandler.PutSocialMediaHandler)
	// 	socialMedia.DELETE("/:id", middlewares.SocialMediaAuthorization(), socialMediaHandler.DeleteSocialMediaHandler)
	// }

	// comment := router.Group("/comment")
	// {
	// 	comment.Use(middlewares.Authentication())
	// 	comment.POST("", commentHandler.PostCommentHandler)
	// 	comment.GET("", commentHandler.GetCommentsHandler)
	// 	comment.GET("/:id", commentHandler.GetCommentHandler)
	// 	comment.PUT("/:id", middlewares.CommentAuthorization(), commentHandler.PutCommentHandler)
	// 	comment.DELETE("/:id", middlewares.CommentAuthorization(), commentHandler.DeleteCommentHandler)
	// }

	return router
}
