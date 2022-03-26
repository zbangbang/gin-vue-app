package main

import (
	"github.com/gin-gonic/gin"
	"zbangbang/gin-vue-app/controller"
	"zbangbang/gin-vue-app/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.RecoveryMiddleware())
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/login", controller.Login)
		userGroup.GET("/info", middleware.AuthMiddleware(), controller.Info)
	}

	categoryGroup := r.Group("/api/category")
	categoryGroup.Use(middleware.AuthMiddleware())
	categoryController := controller.NewCategoryController()
	{
		categoryGroup.POST("", categoryController.Create)
		categoryGroup.PUT("/:id", categoryController.Update)
		categoryGroup.GET("/:id", categoryController.Show)
		categoryGroup.DELETE("/:id", categoryController.Delete)
	}

	postGroup := r.Group("/api/posts")
	postGroup.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	{
		postGroup.POST("", postController.Create)
		postGroup.PUT("/:id", postController.Update)
		postGroup.GET("/:id", postController.Show)
		postGroup.DELETE("/:id", postController.Delete)
		postGroup.GET("/list", postController.PageList)
	}

	return r
}
