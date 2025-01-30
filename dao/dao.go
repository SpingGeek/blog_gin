// Package dao @Author Gopher
// @Date 2025/1/30 11:01:00
// @Desc
package dao

import (
	"blog_gin/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User

	// AddPost 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:wangchen0912@tcp(127.0.0.1:3306)/test_xorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
