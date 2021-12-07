package api

import (
	"fmt"
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"hello/utils/uuid"
	"hello/utils/validator"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var code int

// 发送验证码
func SendvalidateCode(c *gin.Context) {
	var data model.SignUpEmail
	_ = c.ShouldBindJSON(&data)

	code = model.CheckUser(data.EmailName)
	if code == errormsg.SUCCSE {
		uuid := uuid.GetSnowFlakeID()
		code = model.SendvalidateCode(data.EmailName, uuid)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errormsg.GetErrMsg(code),
		},
	)
}

// // 发送验证码
// func CheckvalidateCode(c *gin.Context) {
// 	var data model.SignUpEmail
// 	_ = c.ShouldBindJSON(&data)

// 	code = model.CheckvalidateCode(data.EmailName, data.Uuid)
// 	c.JSON(
// 		http.StatusOK, gin.H{
// 			"status":  code,
// 			"message": errormsg.GetErrMsg(code),
// 		},
// 	)
// }

func CheckCode(c *gin.Context) {
	token := c.Query("token")
	email := model.CheckvalidateCode(token)
	if email != "" {
		code = errormsg.SUCCSE
	} else {
		code = errormsg.ERROR
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errormsg.GetErrMsg(code),
		},
	)
}

func SignUpInClaim(c *gin.Context) {
	var data model.ClaimInfo
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errormsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code = model.CheckCodeAndCreate(data)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errormsg.GetErrMsg(code),
		},
	)
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	username1, _ := c.Get("username")
	username2, _ := username1.(*middleware.MyClaims)
	username3 := username2.Username

	fmt.Print(username3)

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errormsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  username1,
			"data":    data,
			"total":   total,
			"message": errormsg.GetErrMsg(code),
		},
	)
}

// GetUsers 查询用户列表
func GetMe(c *gin.Context) {
	username1, _ := c.Get("username")
	username2, _ := username1.(*middleware.MyClaims)
	username3 := username2.Username

	data, total := model.GetProfile(username3)
	code = errormsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errormsg.GetErrMsg(code),
		},
	)
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errormsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code = model.CheckUser(data.Username)
	if code == errormsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errormsg.ERROR_USERNAME_USED {
		code = errormsg.ERROR_USERNAME_USED
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errormsg.GetErrMsg(code),
		},
	)
}

// Login 后台登陆
func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errormsg.SUCCSE {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errormsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		Id:       user.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "GinBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errormsg.ERROR,
			"message": errormsg.GetErrMsg(errormsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": errormsg.GetErrMsg(200),
		"token":   token,
	})
	return
}
