package model

import (
	"hello/utils/errormsg"

	"github.com/garyburd/redigo/redis"
	"gorm.io/gorm"
)

type Imgcomment struct {
	gorm.Model
	UserId       uint         `json:"user_id"`
	ImgId        uint         `json:"img_id"`
	Content      string       `gorm:"type:varchar(500);not null;" json:"content"`
	Status       int8         `gorm:"type:tinyint;default:2" json:"status"`
	ParentId     uint         `json:"parent_id"`
	Digg         int8         `gorm:"type:tinyint;default:0" json:"digg"`
	DiggArr      []Digg       `gorm:"-"`
	TouserId     uint         `json:"touser_id"`
	FromProfile  Profile      `gorm:"foreignkey:UserId;references:ID"`
	ToProfile    Profile      `gorm:"foreignkey:TouserId;references:ID"`
	CommentId    uint         `gorm:"-" json:"comment_id"`
	Avatar       string       `gorm:"-"`
	CommentChild []Imgcomment `gorm:"-"`
}

func GetImgCommentList(user_id uint, id int, pageSize int, pageNum int) ([]Imgcomment, int64, int64, int) {
	var imgCommentList []Imgcomment
	var count int64
	var floorCount int64

	c := pool.Get()
	defer c.Close()

	db.Find(&Imgcomment{}).Where("img_id = ?", id).Where("status = ?", 1).Count(&count)
	db.Find(&Imgcomment{}).Where("img_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Count(&floorCount)

	db.
		Preload("FromProfile").
		Preload("ToProfile").
		Order("created_at DESC").Where("img_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", 0).Find(&imgCommentList)

	for key, value := range imgCommentList {
		var commentChildList []Imgcomment
		db.
			Preload("FromProfile").
			Preload("ToProfile").
			Order("created_at DESC").Where("img_id = ?", id).Where("status = ?", 1).Where("parent_id = ?", value.ID).Find(&commentChildList)
		num, _ := redis.Int(c.Do("bitcount", imgCommentList[key].ID))

		diggNum, _ := redis.Int(c.Do("getbit", imgCommentList[key].ID, user_id))
		var diggArr Digg
		if diggNum == 1 {
			diggArr.CommentId = imgCommentList[key].ID
			diggArr.ID = user_id
			imgCommentList[key].DiggArr = []Digg{diggArr}
		} else {
			imgCommentList[key].DiggArr = []Digg{}
		}
		imgCommentList[key].Digg = int8(num)

		for childKey, _ := range commentChildList {
			childNum, _ := redis.Int(c.Do("bitcount", commentChildList[childKey].ID))

			diggChildNum, _ := redis.Int(c.Do("getbit", commentChildList[childKey].ID, user_id))
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

		imgCommentList[key].CommentChild = commentChildList
	}

	if err != nil {
		return imgCommentList, 0, 0, errormsg.ERROR
	}
	return imgCommentList, count, floorCount, errormsg.SUCCSE
}

func AddImgComment(ImgComment *Imgcomment) int {
	err = db.Create(&ImgComment).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

func GetParentImgComment(parent_id uint) Imgcomment {
	var imgcomment Imgcomment
	db.Where("id = ?", parent_id).Where("status = ?", 1).First(&imgcomment)
	return imgcomment
}
