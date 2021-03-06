package model

import (
	"hello/utils/errormsg"
	"log"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Imgcontent struct {
	gorm.Model
	Imglist   []*Imglist `gorm:"foreignkey:ContentId;references:ID"`
	Content   string     `gorm:"type:varchar(255)" json:"content"`
	Imgurl    string     `gorm:"type:varchar(255)" json:"imgurl"`
	UserID    uint       `gorm:"type:bigint;not null" json:"user_id"`
	IsComment uint       `gorm:"type:int;not null" json:"is_comment"`
	Profile   Profile    `gorm:"foreignkey:UserID"`
	Digg      uint       `gorm:"-"`
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

func DelImg(id uint) int {
	var imgcontent Imgcontent
	imgcontent.ID = id
	err = db.Preload("Imglist").Delete(&imgcontent).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetImgFavorite(imgFavorite *ImgFavorite) int {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("setbit", "img_like:"+string(strconv.Itoa(int(imgFavorite.ContentId))), imgFavorite.ID, 1)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func SetUnImgFavorite(imgFavorite *ImgFavorite) int {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("setbit", "img_like:"+string(strconv.Itoa(int(imgFavorite.ContentId))), imgFavorite.ID, 0)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func GetImFavorite(imgFavorite *ImgFavorite) int {
	c := pool.Get()
	defer c.Close()
	diggNum, _ := redis.Int(c.Do("getbit", "img_like:"+string(strconv.Itoa(int(imgFavorite.ContentId))), imgFavorite.ID))
	return diggNum
}
