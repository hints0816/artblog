package model

import (
	"hello/utils/errormsg"
	"log"

	"gorm.io/gorm"
)

type Imgcontent struct {
	gorm.Model
	Content string `gorm:"type:varchar(255)" json:"content"`
	Imgurl  string `gorm:"type:varchar(255)" json:"imgurl"`
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
