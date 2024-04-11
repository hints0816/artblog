package model

import (
	"hello/utils/errormsg"
)

type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat    string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat    string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo     string `gorm:"type:varchar(200)" json:"weibo"`
	Bili      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
	Emoji     string `gorm:"type:varchar(100)" json:"emoji"`
	EmojiText string `gorm:"type:varchar(200)" json:"emoji_text"`
}

// GetProfile 获取个人信息设置
func GetProfile(username string) (Profile, int) {
	var profile Profile
	var user User
	db.Where("username = ?", username).Find(&user)
	err = db.Where("ID = ?", user.ID).First(&profile).Error
	if err != nil {
		return profile, errormsg.ERROR
	}
	return profile, errormsg.SUCCSE
}

func GetProfileById(id uint) Profile {
	var profile Profile
	db.Where("ID = ?", id).First(&profile)
	return profile
}

func GetCategoryById(id uint) ([]Category, error) {
    var categories []Category

	subQuery := db.Model(&Article{}).Where("user_id = ?", id).Select("id")
	err := db.Model(&Cateart{}).
	Select("category.id, category.name, category.color, category.text_color, category.icon").
	Distinct().Joins("INNER JOIN category ON cateart.cid = category.id").
	Where("cateart.id IN (?)", subQuery).
	Find(&categories).Error

	if err != nil {
        return nil, err
    }
    return categories, nil
}

func CheckEmail(email string) int {
	var profile Profile
	db.Where("email = ?", email).First(&profile)
	if profile.Email == "" {
		return errormsg.SUCCSE
	} else {
		return errormsg.ERROR
	}
}

// UpdateProfile 更新个人信息设置
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// UpdateProfile 更新个人信息设置
func AddProfile(data *Profile) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}
