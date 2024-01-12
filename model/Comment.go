package model

import (
	"hello/utils/errormsg"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId       uint      `json:"user_id"`
	ArticleId    uint      `json:"article_id"`
	Content      string    `gorm:"type:varchar(500);not null;" json:"content"`
	Status       int8      `gorm:"type:tinyint;default:2" json:"status"`
	ParentId     uint      `json:"parent_id"`
	Digg         int8      `gorm:"type:tinyint;default:0" json:"digg"`
	DiggArr      []Digg    `gorm:"-"`
	TouserId     uint      `json:"touser_id"`
	FromProfile  Profile   `gorm:"foreignkey:UserId;references:ID"`
	ToProfile    Profile   `gorm:"foreignkey:TouserId;references:ID"`
	CommentId    uint      `gorm:"-" json:"comment_id"`
	Avatar       string    `gorm:"-"`
	Title        string    `gorm:"-"`
	CommentChild []Comment `gorm:"-"`
}

type Digg struct {
	ID        uint `gorm:"primarykey"`
	CommentId uint `gorm:"primarykey" json:"comment_id"`
}

func GetCommentList(user_id uint, id int, pageSize int, pageNum int) ([]Comment, int64, int64, int) {
	var commentList []Comment
	var count int64
	var floorCount int64

	c := pool.Get()
	defer c.Close()

	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&count)
	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Count(&floorCount)

	db.
		Preload("FromProfile").
		Preload("ToProfile").
		Order("created_at DESC").Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Find(&commentList)

	for key, value := range commentList {
		var commentChildList []Comment
		db.
			Preload("FromProfile").
			Preload("ToProfile").
			Order("created_at DESC").Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", value.ID).Find(&commentChildList)
		num, _ := redis.Int(c.Do("bitcount", "comment_like:"+string(strconv.Itoa(int(commentList[key].ID)))))

		diggNum, _ := redis.Int(c.Do("getbit", "comment_like:"+string(strconv.Itoa(int(commentList[key].ID))), user_id))
		var diggArr Digg
		if diggNum == 1 {
			diggArr.CommentId = commentList[key].ID
			diggArr.ID = user_id
			commentList[key].DiggArr = []Digg{diggArr}
		} else {
			commentList[key].DiggArr = []Digg{}
		}
		commentList[key].Digg = int8(num)

		for childKey, _ := range commentChildList {
			childNum, _ := redis.Int(c.Do("bitcount", "comment_like:"+string(strconv.Itoa(int(commentChildList[childKey].ID)))))

			diggChildNum, _ := redis.Int(c.Do("getbit", "comment_like:"+string(strconv.Itoa(int(commentChildList[childKey].ID))), user_id))
			var diggChildArr Digg
			if diggChildNum == 1 {
				diggChildArr.CommentId = commentChildList[childKey].ID
				diggChildArr.ID = user_id
				commentChildList[childKey].DiggArr = []Digg{diggChildArr}
			} else {
				commentChildList[childKey].DiggArr = []Digg{}
			}
			commentChildList[childKey].Digg = int8(childNum)
		}

		commentList[key].CommentChild = commentChildList
	}

	if err != nil {
		return commentList, 0, 0, errormsg.ERROR
	}
	return commentList, count, floorCount, errormsg.SUCCSE
}
# Global logging configuration
log4j.rootLogger= debug, console, 
 # SqlMap logging configuration...

log4j.logger.java.sql.ResultSet=DEBUG
log4j.logger.org.apache=INFO
log4j.logger.java.sql.Connection=DEBUG
log4j.logger.java.sql.Statement=DEBUG
log4j.logger.java.sql.PreparedStatement=DEBUG

log4j.logger.org.springframework=DEBUG
log4j.logger.com.mchange=INFO
log4j.logger.com.gree=DEBUG
log4j.logger.org.jbpm=INFO

# console output... 
log4j.appender.console=org.apache.log4j.ConsoleAppender
log4j.appender.console.layout=org.apache.log4j.PatternLayout
log4j.appender.console.layout.ConversionPattern=[%p][%F:%L]:%m%n

# File output
log4j.appender.f=org.apache.log4j.DailyRollingFileAppender
log4j.appender.f.Threshold=INFO
log4j.appender.f.File=${catalina.base}/logs_cp/log4j.log
log4j.appender.f.DatePattern=yyyy-MM-dd'.log' 
log4j.appender.f.layout=org.apache.log4j.PatternLayout
log4j.appender.f.layout.ConversionPattern=[%p][%d][%F\:%L]\:%m%n
func GetCommentListAdmin(user_id int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var count int64

	c := pool.Get()
	defer c.Close()

	db.Where("user_id = ?", user_id).Or("touser_id = ?", user_id).Find(&Comment{}).Count(&count)

	err := db.
		Preload("FromProfile").
		Preload("ToProfile").
		Order("created_at DESC").Where("user_id = ?", user_id).Or("touser_id = ?", user_id).Find(&commentList).Error

	for childKey, _ := range commentList {
		var art Article
		articleId := commentList[childKey].ArticleId
		db.Where("id =?", articleId).Find(&art)
		commentList[childKey].Title = art.Title
	}

	if err != nil {
		return commentList, 0, errormsg.ERROR
	}
	return commentList, count, errormsg.SUCCSE
}

// AddComment 新增主评论
func AddComment(comment *Comment) int {
	err = db.Create(&comment).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// AddComment 新增主评论
func DelComment(id uint) int {
	var comment Comment
	comment.ID = id
	err = db.Delete(&comment).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetDigg(digg *Digg) int {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("setbit", "comment_like:"+string(strconv.Itoa(int(digg.CommentId))), digg.ID, 1)
	// 使用redis代替mysql
	// err = db.Create(&digg).Error
	// var comment Comment
	// db.Model(&comment).Where("id = ?", digg.CommentId).UpdateColumn("digg", gorm.Expr("digg + ?", 1))
	if err != nil {
		fmt.Println(err)
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetUnDigg(digg *Digg) int {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("setbit", "comment_like:"+string(strconv.Itoa(int(digg.CommentId))), digg.ID, 0)
	// 使用redis代替mysql
	// var comment Comment
	// db.Model(&comment).Where("id = ?", digg.CommentId).UpdateColumn("digg", gorm.Expr("digg - ?", 1))
	if err != nil {
		fmt.Println(err)
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func GetParentComment(parent_id uint) Comment {
	var comment Comment
	db.Where("id = ?", parent_id).Where("status = ?", 1).First(&comment)
	return comment
}
