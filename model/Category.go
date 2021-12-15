package model

import (
	"hello/utils/errormsg"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Color     string `gorm:"type:varchar(255);not null" json:"color"`
	TextColor string `gorm:"type:varchar(255);not null" json:"text_color"`
	Icon      string `gorm:"type:varchar(255);not null" json:"icon"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errormsg.ERROR_CATENAME_USED //2001
	}
	return errormsg.SUCCSE
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errormsg.SUCCSE
}

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	// var maps = make(map[string]interface{})
	// maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Updates(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCSE
}
