package model

import (
	"hello/utils/errormsg"
	"log"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Imgcontent struct {
	gorm.Model
	Imglist []*Imglist `gorm:"foreignkey:ContentId;references:ID"`
	Content string     `gorm:"type:varchar(255)" json:"content"`
	Imgurl  string     `gorm:"type:varchar(255)" json:"imgurl"`
	UserID  uint       `gorm:"type:bigint;not null" json:"user_id"`
	Profile Profile    `gorm:"foreignkey:UserID"`
}

type Imglist struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	ImgUrl    string    `gorm:"type:varchar(255)" json:"img_url"`
	ContentId uint      `gorm:"type:int" json:"content_id"`
	Imgtag    []*Imgtag `gorm:"foreignkey:ID;references:ID"`
}

type Imgtag struct {
	ID   int             `gorm:"type:int" json:"id"`
	Top  decimal.Decimal `gorm:"type:float64" json:"top"`
	Left decimal.Decimal `gorm:"type:float64" json:"left"`
	Tag  string          `gorm:"type:varchar(255)" json:"tag"`
}

func ListImage(pageSize int, pageNum int) ([]Imgcontent, int, int64) {
	var imgContent []Imgcontent
	var total int64

	tx := db.Order("created_at DESC")
	err = tx.Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&imgContent).Error

	tx2 := db.Model(&imgContent)
	tx2.Count(&total)

	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return imgContent, errormsg.SUCCSE, total
}

func PostArt(data *Imgcontent) int {
	err := db.Create(&data).Error
	if err != nil {
		log.Fatalln(err)
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

func GetImgInfo(id int) (Imgcontent, int) {
	var imgcontent Imgcontent
	db.Preload("Imglist.Imgtag").Preload("Imglist").Preload("Profile").Where("id =?", id).First(&imgcontent)
	return imgcontent, errormsg.SUCCSE
}
