package model

import (
	"hello/utils/errormsg"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId       uint      `json:"user_id"`
	Profile      *Profile  `gorm:"foreignkey:UserId;references:ID"`
	ArticleId    uint      `json:"article_id"`
	Username     string    `json:"username"`
	Content      string    `gorm:"type:varchar(500);not null;" json:"content"`
	Status       int8      `gorm:"type:tinyint;default:2" json:"status"`
	ParentId     int8      `json:"parent_id"`
	Digg         int8      `gorm:"type:tinyint;default:2" json:"digg"`
	ToId         int8      `json:"to_id"`
	CommentChild []Comment `gorm:"-"`
}

type Commentchild struct {
	gorm.Model
	FromProfile *Profile `gorm:"foreignkey:From"`
	From        uint     `json:"from"`
	ToProfile   *Profile `gorm:"foreignkey:To"`
	To          uint     `json:"to"`
	Cid         uint     `gorm:"primary_key;not null" json:"cid"`
	Content     string   `gorm:"type:varchar(500);not null;" json:"content"`
	ArticleId   uint     `json:"article_id"`
}

type CommentchildVo struct {
	ArticleId uint   `json:"article_id"`
	Content   string `json:"content"`
	ParentId  uint   `json:"parent_id"`
}

func GetRecurComment(commentList []Comment, id int) {
	for key, value := range commentList {
		var commentChildList []Comment
		db.Order("created_at DESC").Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", value.ID).Find(&commentChildList)

		copy(commentList[key:], commentChildList)

		commentList[key].CommentChild = commentChildList

		GetRecurComment(commentList[key].CommentChild, id)
	}
}

func GetCommentListFront(id int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	var total1 int64

	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)

	db.Order("created_at DESC").Where("article_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Find(&commentList)

	GetRecurComment(commentList, id)

	// db.Preload("Commentchild", func(db *gorm.DB) *gorm.DB {
	// 	return db.Order("commentchild.created_at DESC")
	// }).Preload("Profile").Preload("Commentchild.FromProfile").Preload("Commentchild.ToProfile").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("article_id = ?", id).Find(&commentList)

	if err != nil {
		return commentList, 0, errormsg.ERROR
	}
	return commentList, total + total1, errormsg.SUCCSE
}

// AddComment 新增主评论
func AddComment(data *Comment) int {
	err = db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// AddCommentChild 新增子评论
func AddCommentChild(data *Commentchild, commentchildVo *CommentchildVo) int {
	var commentchild Commentchild
	db.Where("id = ?", commentchildVo.ParentId).First(&commentchild)
	data.Content = commentchildVo.Content
	data.To = commentchild.From
	data.Cid = commentchild.Cid
	data.ArticleId = commentchildVo.ArticleId

	err = db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}
