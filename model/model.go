// Package model @Author Gopher
// @Date 2025/1/30 10:59:00
// @Desc
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
