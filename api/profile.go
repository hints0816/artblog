package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"hello/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 查询单个用户
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetProfileById(uint(id))
	c.JSON(
		http.StatusOK, gin.H{
			"status": 200,
			"data":   data,
		},
	)
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	_ = c.ShouldBindJSON(&data)
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	dataUpdate := model.GetProfileById(userinfo.Id)
	dataUpdate.Name = data.Name
	dataUpdate.Wechat = data.Wechat
	dataUpdate.Qqchat = data.Qqchat
	dataUpdate.Weibo = data.Weibo
	dataUpdate.Bili = data.Bili
	code := model.UpdateProfile(dataUpdate.ID, &dataUpdate)
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   data,
		},
	)
}

func UploadAvatarImage(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	fileSize := fileHeader.Size
	contentType := fileHeader.Header.Get("Content-Type")
	url, code := model.UpLoadFile(file, contentType, fileSize)

	profile := model.GetProfileById(userinfo.Id)
	profile.Avatar = url
	code1 := model.UpdateProfile(profile.ID, &profile)

	c.JSON(http.StatusOK, gin.H{
		"status":  code1,
		"message": errormsg.GetErrMsg(code),
	})
}

func EditEmoji(c *gin.Context) {
	var data model.Profile
	_ = c.ShouldBindJSON(&data)
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	profile := model.GetProfileById(userinfo.Id)
	profile.Emoji = data.Emoji
	profile.EmojiText = data.EmojiText
	code := model.UpdateProfile(profile.ID, &profile)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func GetWebmasterInfo(c *gin.Context) {
	webMaster := utils.WebMaster
	id, _ := strconv.Atoi(webMaster)
	data := model.GetProfileById(uint(id))
	data1,_ := model.GetCategoryById(uint(id))
	c.JSON(
		http.StatusOK, gin.H{
			"status": 200,
			"data":   data,
			"data1":   data1,
		},
	)
}