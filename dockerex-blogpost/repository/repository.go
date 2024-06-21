package repository

import (
	"blogpost/adapter"
	"blogpost/helper"
	"blogpost/models"
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Database struct {
	Dbcon *gorm.DB
	FuncCall
}

type FuncCall interface {
	ToValidateUserLogin(loginReq *models.Logincredentials) error
	UserCreation(newUser models.Logincredentials) error
	ToUpdataUserData(userdata models.Logincredentials, id int) error
	ToDeleteUserData(id int) error
	ToCheckCategories(id int) error
	ToCreatePost(blog models.Posts) error
	ToUpdatePost(blog models.Posts, id int) error
	ToDeletePost(id int) error
	ToCreateCategory(cateory models.Categories) error
	ToUpdateCategory(category models.Categories, id int) error
	ToDeleteCategory(id int) error
	ToCreateNewComment(comment models.Comments) error
	ToUpdateComment(comment models.Comments, id int) error
	ToDeleteComments(id uint64) error
	ToGetCommentsByCommentId(id int) ([]models.CommentsResponse, error)
	ToGetPostByCategory(categoryname string) ([]models.PostsResponse, error)
	ToGetPostsOverview() (interface{}, error)
	ToSetRedisCache(token, uuid string) error
	ToGetPostComment(id int) ([]models.Comments, error)
	ToGetCategoryByID(id int) (name string, err error)
}

var Dbn Database

func ToSetDB(db *gorm.DB) {
	Dbn = Database{
		Dbcon: db,
	}
	// return &Dbn
}

func (db *Database) ToSignUpUser(newUser *models.Logincredentials) error {
	usercreation := new(models.Logincredentials)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usercreation.Username = newUser.Username
	usercreation.Password = string(hashedPassword)
	usercreation.Role = newUser.Role
	if err := db.Dbcon.Create(&usercreation).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}
func (db *Database) ToLoginUser(loginReq *models.Logincredentials) (models.Logincredentials, error) {
	var user models.Logincredentials
	if err := db.Dbcon.Where("user_name = ?", loginReq.Username).First(&user).Error; err != nil {
		return models.Logincredentials{}, errors.New("cannot find the user")
	}
	//migrated to middleware

	// Dbn.ToSetRedisCache(token)
	return user, nil
}

func (db *Database) ToUpdataUserData(userdata models.Logincredentials, id int) error {
	existingContent := new(models.Logincredentials)
	if err := db.Dbcon.Where("id=?", id).Find(&existingContent).Error; err != nil {
		return err
	}
	if userdata.Username != "" && userdata.Password == "" {
		fmt.Print("username setting")
		existingContent.Username = userdata.Username
	}
	if userdata.Username == "" && userdata.Password != "" {
		fmt.Print("passwords setting")
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userdata.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingContent.Password = string(hashedPassword)
	}
	if err := db.Dbcon.Debug().Save(&existingContent).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) ToDeleteUserData(id int) error {
	deleteUser := models.Logincredentials{}
	var err *gorm.DB
	if err = db.Dbcon.Model(&deleteUser).Where("Id = ?", id).Delete(&deleteUser); err.Error != nil {
		return fmt.Errorf("post does not exits")
	}
	if err.RowsAffected == 0 {
		return errors.New("cannot delete post")
	}
	return nil
}

func (db *Database) ToCreatePost(blog models.Posts) error {
	if err := db.Dbcon.Create(&blog).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) ToUpdatePost(blog models.Posts, id int) error {
	if err := db.Dbcon.Where("id=?", id).UpdateColumns(&blog).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) ToCheckCategories(id int) error {
	category := models.Categories{}
	var err *gorm.DB
	if err = db.Dbcon.Where("id=?", id).Find(&category); err.Error != nil {
		return fmt.Errorf("categoy not found,%v", err)
	}
	if err.RowsAffected == 0 {
		return errors.New("categoy not found")
	}
	return nil
}
func (db *Database) ToDeletePost(id int) error {
	deletePost := models.Posts{}
	var err *gorm.DB
	if err = db.Dbcon.Model(&deletePost).Unscoped().Where("Id = ?", id).Delete(&deletePost); err.Error != nil {
		return fmt.Errorf("post does not exits")
	}
	if err.RowsAffected == 0 {
		return errors.New("cannot delete post")
	}
	return nil
}

func (db *Database) ToGetAllPosts() ([]models.Posts, error) {
	var postsResponse []models.Posts
	if err := db.Dbcon.Model(&models.Posts{}).Find(&postsResponse).Error; err != nil {
		return nil, err
	}
	return postsResponse, nil
}
func (db *Database) ToGetPostComment(id int) ([]models.CommentsResponse, error) {
	var postsResponse []models.CommentsResponse
	if err := db.Dbcon.Model(models.Comments{}).Select("commentcontent").Where("id=?", id).Scan(&postsResponse).Error; err != nil {
		return nil, err
	}
	return postsResponse, nil
}
func (db *Database) ToGetCategoryByID(id int) (name string, err error) {
	if err = db.Dbcon.Model(&models.Categories{}).Select("name").Where("id=?", id).Find(&name).Error; err != nil {
		return
	}
	return
}
func (db *Database) ToGetCategoryByName(name string) (categoryid int, err error) {
	if err = db.Dbcon.Model(&models.Categories{}).Select("id").Where("name=?", name).Find(&categoryid).Error; err != nil {
		return
	}
	return
}
func (db *Database) ToGetPostByCategory(id int) ([]models.Posts, error) {
	var postsbycategory []models.Posts
	if err := db.Dbcon.Model(&models.Posts{}).Where("?=ANY(categoriesid)", id).Find(&postsbycategory).Error; err != nil {
		return nil, err
	}
	return postsbycategory, nil
}

func (db *Database) ToCreateNewComment(comment models.Comments) error {
	if err := db.Dbcon.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) ToUpdateComment(comment models.Comments, id int) error {
	var err *gorm.DB
	if err = db.Dbcon.Where("id=?", id).UpdateColumns(&comment); err.Error != nil {
		return fmt.Errorf("cannot updated comment")
	}
	if err.RowsAffected == 0 {
		return fmt.Errorf("cannot updated comment")
	}
	return nil
}

func (db *Database) ToDeleteComments(id uint64) error {
	comment := models.Comments{}
	if err := db.Dbcon.Where("Id = ?", id).First(&comment).Error; err != nil {
		return fmt.Errorf("comment does not exits")
	}
	if err := db.Dbcon.Where("Id = ?", id).Delete(&comment).Error; err != nil {
		return fmt.Errorf("comment does not exits")
	}
	return nil
}
func (db *Database) ToGetCommentsByPostId(id int) ([]models.CommentsOnPosts, error) {
	commentresponse := []models.CommentsOnPosts{}
	if err := db.Dbcon.Table("posts p").Select("p.title", "p.content", "p.excerpt", "c.commentcontent").Joins("JOIN comments c ON p.id = c.user_id").Where("p.id=?", id).Scan(&commentresponse).Error; err != nil {
		return nil, err
	}
	return commentresponse, nil
}
func (db *Database) ToGetCommentsByCommentId(id int) ([]models.CommentsResponse, error) {
	commentresponse := []models.CommentsResponse{}
	if err := db.Dbcon.Table("comments c").Select("c.commentcontent", "l.user_name").Joins("join logincredentials l on c.user_id=l.id").Where("c.id=?", id).Scan(&commentresponse).Error; err != nil {
		return nil, err
	}
	return commentresponse, nil
}
func (db *Database) ToGetCommentsByUserId(id int) ([]models.CommentsResponse, error) {
	commentresponse := []models.CommentsResponse{}
	if err := db.Dbcon.Table("comments c").Select("c.commentcontent", "l.user_name").Joins("join logincredentials l on c.user_id=l.id").Where("l.id=?", id).Scan(&commentresponse).Error; err != nil {
		return nil, err
	}
	return commentresponse, nil
}

func (db *Database) ToCreateCategory(cateory models.Categories) error {
	if err := db.Dbcon.Create(&cateory).Error; err != nil {
		return err
	}
	return nil
}
func (db *Database) ToDeleteCategory(id int) error {
	deleteCategories := models.Categories{}
	if err := db.Dbcon.Model(&deleteCategories).Unscoped().Where("Id = ?", id).Delete(&deleteCategories).Error; err != nil {
		return fmt.Errorf("post does not exits")
	}
	return nil
}

func (db *Database) ToUpdateCategory(category models.Categories, id int) error {
	if err := db.Dbcon.Where("id=?", id).UpdateColumns(&category).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) ToGetPostsByArchieves(year int, month time.Month) ([]models.ArcivedPostsResponse, error) {
	var postsResponse []models.ArcivedPostsResponse

	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	err := db.Dbcon.Table("posts p").Select("p.title", "p.content", "p.excerpt", "c.commentcontent", "cat.name AS category_name").Joins("JOIN comments c ON p.id = c.user_id").Joins("JOIN categories cat ON cat.id = ANY(p.categoriesid)").Where("p.created_at >= ? AND p.created_at <= ?", firstOfMonth, lastOfMonth).Scan(&postsResponse).Error
	if err != nil {
		return nil, err
	}
	return postsResponse, nil
}

func (db *Database) ToGetPostsOverview() (interface{}, error) {
	overview := models.OverviewData{}
	result := db.Dbcon.Table("posts p").Select("COUNT(p.id) AS no_of_posts", "COUNT(c.id) AS no_of_comments").Joins("JOIN comments c ON p.id = c.post_id").Scan(&overview)
	if result.Error != nil {
		return nil, result.Error
	}
	var date time.Time
	db.Dbcon.Table("posts").Select("created_at").Scan(&date)
	y, m, d := helper.DateDifference(date, time.Now())
	dates := fmt.Sprintf("%d:Y %d:M %d:D", y, m, d)
	overview.PublishedOn = dates
	return overview, nil
}

func (db *Database) ToSetRedisCache(userdata models.Logincredentials, uuid string) error {
	fmt.Print(uuid)
	if err := adapter.RedisConnection().HMSet(context.TODO(), uuid, "id", userdata.ID, "role", userdata.Role).Err(); err != nil {
		return err
	}

	return nil
}
func (db *Database) ToGetRedisCache(uuid string) (string, error) {
	fmt.Print("Getting redis cache")
	result, err := adapter.RedisConnection().HGet(context.TODO(), uuid, "role").Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func (db *Database) ToGetRedisUserID(uuid string) (string, error) {
	result, err := adapter.RedisConnection().HGet(context.TODO(), uuid, "id").Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
func (db *Database) ToDeleteRedisCache(uuid string) error {
	if err := adapter.RedisConnection().Del(context.TODO(), uuid).Err(); err != nil {
		return err
	}
	return nil
}
