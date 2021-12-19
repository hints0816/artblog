package model

import (
	"hello/utils/errormsg"

	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" label:"角色码"`
	Github   int    `gorm:"type:int" json:"github"`
}

type ClaimInfo struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errormsg.ERROR_USERNAME_USED //1001
	}
	return errormsg.SUCCSE
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}
	return errormsg.SUCCSE
}

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errormsg.ERROR
	}
	return user, errormsg.SUCCSE
}

// GetUser 查询用户
func GetUserByGithub(github int) (User, int) {
	var user User
	err := db.Limit(1).Where("github = ?", github).Find(&user).Error
	if err != nil {
		return user, errormsg.ERROR
	}
	return user, errormsg.SUCCSE
}

func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,username,role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errormsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errormsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errormsg.ERROR_USER_NO_RIGHT
	}
	return user, errormsg.SUCCSE
}
