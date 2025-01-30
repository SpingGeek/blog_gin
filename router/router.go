// Package router @Author Gopher
// @Date 2025/1/30 10:58:00
// @Desc
package router

import (
	"blog_gin/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/", controller.Index)
	e.POST("/register", controller.RegisterUser)
	e.GET("/register", controller.GoRegister)
	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/post_detail", controller.PostDetail)
	e.Run(":8081")
}
