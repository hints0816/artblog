package model

import (
	"hello/utils/errormsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	db.Where("id = ?", id).First(&art)
	return art, errormsg.SUCCSE
}

func ListArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var total int64

	db.Find(&artList)
	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artList, errormsg.SUCCSE, total
}
