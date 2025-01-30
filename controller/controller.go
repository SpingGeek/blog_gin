// Package controller @Author Gopher
// @Date 2025/1/30 10:59:00
// @Desc
package controller

import (
	"blog_gin/dao"
	"blog_gin/model"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)

	c.Redirect(200, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	// content := blackfriday.Run([]byte(p.Content))
	md := []byte(p.Content)
	md = markdown.NormalizeNewlines(md)
	html := markdown.ToHTML(md, parser, nil)

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(html),
	})
}
