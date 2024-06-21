package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model   `json:"-"`
	Title        string        `gorm:"column:title" validate:"required" json:"title"`
	Content      string        `gorm:"column:content" validate:"required" json:"content"`
	Status       string        `gorm:"column:status"  validate:"required" json:"status"`
	Excerpt      string        `gorm:"column:excerpt" json:"excerpt"`
	Categoriesid pq.Int64Array `gorm:"type:integer[];column:categoriesid" validate:"required" json:"categoriesid"`
}
type Categories struct {
	Id          int    `gorm:"primarykey;autoIncrement;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

type Comments struct {
	Id               int              `gorm:"primarykey;autoIncrement;column:id" json:"id,omitempty"`
	PostId           int              `gorm:"column:post_id" json:",omitempty"`
	Posts            Posts            `gorm:"foreignKey:PostId" json:",omitempty"`
	UserId           int              `gorm:"column:user_id" json:",omitempty"`
	Logincredentials Logincredentials `gorm:"foreignKey:UserId" json:"user"`
	CommentContent   string           `gorm:"column:commentcontent" json:"commentcontent"`
}

type Logincredentials struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"column:user_name" json:"username,omitempty"`
	Password   string `gorm:"column:password" json:"password,omitempty"`
	Role       string `gorm:"column:role" json:"role,omitempty"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type CommentsResponse struct {
	CommentContent string `gorm:"column:commentcontent" json:"commentcontent"`
	User_name      string `gorm:"column:user_name" json:"username,omitempty"`
}

type CommentsOnPosts struct {
	Title          string `gorm:"column:title" validate:"required" json:"title"`
	Content        string `gorm:"column:content" validate:"required" json:"content"`
	Excerpt        string `gorm:"column:excerpt" json:"excerpt"`
	CommentContent string `gorm:"column:commentcontent" json:"commentcontent"`
}

type PostsResponse struct {
	Title          string             `json:"title,omitempty"`
	Content        string             `json:"content,omitempty"`
	Excerpt        string             `json:"excerpt,omitempty"`
	CommentContent []CommentsResponse `gorm:"column:commentcontent" json:"commentcontent,omitempty"`
	CategoryName   []string           `json:"category_name,omitempty"`
}
type ArcivedPostsResponse struct {
	Title          string `json:"title,omitempty"`
	Content        string `json:"content,omitempty"`
	Excerpt        string `json:"excerpt,omitempty"`
	CommentContent string `gorm:"column:commentcontent" json:"commentcontent,omitempty"`
	CategoryName   string `json:"category_name,omitempty"`
}
type OverviewData struct {
	NoOfPosts    int    `json:"no_of_posts"`
	NoOfComments int    `json:"no_of_comments"`
	PublishedOn  string `json:"published_on"`
}
