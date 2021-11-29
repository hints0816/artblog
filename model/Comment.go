package model

import (
	"hello/utils/errormsg"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId       uint      `json:"user_id"`
	ArticleId    uint      `json:"article_id"`
	Username     string    `json:"username"`
	Content      string    `gorm:"type:varchar(500);not null;" json:"content"`
	Status       int8      `gorm:"type:tinyint;default:2" json:"status"`
	ParentId     uint      `json:"parent_id"`
	Digg         int8      `gorm:"type:tinyint;default:0" json:"digg"`
	DiggArr      []Digg    `gorm:"foreignkey:CommentId;references:ID"`
	ToId         uint      `json:"to_id"`
	CommentId    uint      `gorm:"-" json:"comment_id"`
	ToUsername   string    `json:"to_username"`
	Avatar       string    `json:"avatar"`
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

	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&count)
	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Count(&floorCount)

	db.Preload("DiggArr", "id =?", user_id).Order("created_at DESC").Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Find(&commentList)

	for key, value := range commentList {
		var commentChildList []Comment
		db.Preload("DiggArr", "id =?", user_id).Order("created_at DESC").Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", value.ID).Find(&commentChildList)
		commentList[key].CommentChild = commentChildList
	}

	if err != nil {
		return commentList, 0, 0, errormsg.ERROR
	}
	return commentList, count, floorCount, errormsg.SUCCSE
}

// AddComment 新增主评论
func AddComment(comment *Comment) int {
	err = db.Create(&comment).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetDigg(digg *Digg) int {
	err = db.Create(&digg).Error
	var comment Comment
	db.Model(&comment).Where("id = ?", digg.CommentId).UpdateColumn("digg", gorm.Expr("digg + ?", 1))
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetUnDigg(digg *Digg) int {
	err = db.Delete(&digg).Error
	var comment Comment
	db.Model(&comment).Where("id = ?", digg.CommentId).UpdateColumn("digg", gorm.Expr("digg - ?", 1))
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func GetParentComment(parent_id uint) Comment {
	var comment Comment
	db.Where("id = ?", parent_id).Where("status = ?", 1).First(&comment)
	return comment
}
