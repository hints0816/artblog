package model

import (
	"hello/utils/errormsg"
	"log"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Cateart      []*Cateart `gorm:"foreignkey:ID;"`
	Title        string     `gorm:"type:varchar(100);not null" json:"title"`
	Desc         string     `gorm:"type:varchar(200)" json:"desc"`
	Content      string     `gorm:"type:longtext" json:"content"`
	Img          string     `gorm:"type:varchar(100)" json:"img"`
	CommentCount int        `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int        `gorm:"type:int;not null;default:0" json:"read_count"`
	Status       int8       `gorm:"type:tinyint;default:0" json:"status"`
	UserID       uint       `gorm:"type:bigint;not null" json:"user_id"`
}

type Cateart struct {
	Category Category `gorm:"foreignkey:Cid;"`
	ID       int      `gorm:"primary_key;not null" json:"id"`
	Cid      int      `gorm:"primary_key;not null" json:"cid"`
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		log.Fatalln(err)
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

// CreateArt 更新文章
func UpdateArt(data *Article) int {
	err = db.Model(&Article{}).Where("ID = ?", data.ID).Updates(&data).Error
	if err != nil {
		log.Fatalln(err)
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

func DelArt(id uint) int {
	var article Article
	article.ID = id
	err = db.Delete(&article).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	db.Where("id =?", id).First(&art)
	return art, errormsg.SUCCSE
}

func ListArticle(userId int, cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var ids []string
	var total int64

	if cid != 0 {
		db.Model(&Cateart{}).Where("cid =?", cid).Pluck("id", &ids)
	}
	tx := db.
		Order("created_at DESC").
		Preload("Cateart").
		Preload("Cateart.Category").
		Where("status =?", 1)
	if cid != 0 {
		tx = tx.Where("id in ?", ids)
	}
	if userId != 0 {
		tx = tx.Where("user_id = ?", userId)
	}
	err = tx.Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&artList).Error

	tx2 := db.Model(&artList).Where("status =?", 1)
	if cid != 0 {
		tx2 = tx2.Where("id in ?", ids)
	}
	if userId != 0 {
		tx2 = tx2.Where("user_id = ?", userId)
	}
	tx2.Count(&total)

	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artList, errormsg.SUCCSE, total
}

func ListArticleRepository(userId uint, id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var ids []string
	var total int64

	if id != 0 {
		db.Model(&Cateart{}).Where("cid =?", id).Pluck("id", &ids)
	}
	tx := db.
		Order("created_at DESC").
		Preload("Cateart").
		Preload("Cateart.Category").
		Where("user_id =?", userId)
	if id != 0 {
		tx = tx.Where("id in ?", ids)
	}
	err = tx.Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&artList).Error
	if id != 0 {
		db.Model(&artList).Where("user_id =?", userId).Where("id in ?", ids).Count(&total)
	} else {
		db.Model(&artList).Where("user_id =?", userId).Count(&total)
	}

	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return artList, errormsg.SUCCSE, total
}
