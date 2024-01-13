package routes

import (
	_ "github.com/ariwiraa/my-gram/docs"
	"github.com/ariwiraa/my-gram/handler"
	"github.com/ariwiraa/my-gram/middlewares"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

type RouterHandler struct {
	AuthHandler       handler.AuthHandler
	PhotoHandler      handler.PhotoHandler
	CommentHandler    handler.CommentHandler
	LikesHandler      handler.UserLikesPhotosHandler
	FollowsHandler    handler.FollowHandler
	UserHandler       handler.UserHandler
	UploadFileHandler handler.UploadFileHandler
}

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
func NewRouter(routerHandler RouterHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", routerHandler.AuthHandler.PostUserRegisterHandler)
	router.GET("/verify-email", routerHandler.AuthHandler.VerifyEmail)
	router.POST("/resend-email", routerHandler.AuthHandler.ResendEmail)
	router.POST("/signin", routerHandler.AuthHandler.PostUserLoginHandler)
	router.PUT("/refresh", routerHandler.AuthHandler.PutAccessTokenHandler)
	router.DELETE("/signout", middlewares.Authentication(), routerHandler.AuthHandler.LogoutHandler)
	router.POST("/files/upload", middlewares.Authentication(), routerHandler.UploadFileHandler.UploadFileHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	photo := router.Group("/photos")
	{
		// Photo
		photo.Use(middlewares.Authentication())
		photo.POST("", routerHandler.PhotoHandler.PostPhotoHandler)
		photo.GET("/all", routerHandler.PhotoHandler.GetPhotosHandler)
		photo.GET("", routerHandler.PhotoHandler.GetPhotosByUserIdHandler)
		photo.GET("/:id", routerHandler.PhotoHandler.GetPhotoHandler)
		photo.PUT("/:id", routerHandler.PhotoHandler.PutPhotoHandler)
		photo.DELETE("/:id", routerHandler.PhotoHandler.DeletePhotoHandler)

		// Likes Photo
		photo.POST("/:id/likes", routerHandler.LikesHandler.PostLikesHandler)
		photo.GET("/:id/likes", routerHandler.LikesHandler.GetUsersWhoLikedPhotosHandler)

		// Comments
		photo.POST("/:id/comments", routerHandler.CommentHandler.PostCommentHandler)
		photo.GET("/:id/comments", routerHandler.CommentHandler.GetCommentsHandler)
		photo.GET("/:id/comments/:commentId", routerHandler.CommentHandler.GetCommentHandler)
		photo.PUT("/:id/comments/:commentId", routerHandler.CommentHandler.PutCommentHandler)
		photo.DELETE("/:id/comments/:commentId", routerHandler.CommentHandler.DeleteCommentHandler)
	}

	me := router.Group("/me")
	{
		me.Use(middlewares.Authentication())
		me.GET("/liked/photos", routerHandler.LikesHandler.GetPhotosLikedHandler)
	}

	users := router.Group("/users")
	{
		users.Use(middlewares.Authentication())
		// Follow
		users.POST("/:id/follows", routerHandler.FollowsHandler.PostFollowHandler)
		users.GET("/:username/followers", routerHandler.FollowsHandler.GetFollowersHandler)
		users.GET("/:username/followings", routerHandler.FollowsHandler.GetFollowingsHandler)

		// Profile
		users.GET("/profile/:username", routerHandler.UserHandler.GetUserProfileHandler)
	}

	return router
}
