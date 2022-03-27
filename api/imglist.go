package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"hello/utils/img"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListImage(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 15
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.ListImage(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

func PostImg(c *gin.Context) {
	var data model.Imgcontent

	// file, fileHeader, _ := c.Request.FormFile("file")
	// content := c.Request.FormValue("content")
	// fileSize := fileHeader.Size
	// contentType := fileHeader.Header.Get("Content-Type")
	// url, code := model.UpLoadFile(file, contentType, fileSize)

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&data)
	data.UserID = userinfo.Id
	image := img.ResizeImg(data.Imgurl)
	url, _ := model.UpLoadFileImg(image)
	data.Imgurl = url
	code = model.PostArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func GetImgInfo(c *gin.Context) {

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetImgInfo(id)
	if userinfo != nil {
		var imgFavorite model.ImgFavorite
		imgFavorite.ContentId = uint(id)
		imgFavorite.ID = userinfo.Id
		digg := model.GetImFavorite(&imgFavorite)
		data.Digg = uint(digg)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

func DelImgInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelImg(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func SetFavorite(c *gin.Context) {
	var favorite model.ImgFavorite
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&favorite)
	favorite.ID = userinfo.Id
	code = model.SetImgFavorite(&favorite)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func SetUnFavorite(c *gin.Context) {
	var favorite model.ImgFavorite
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&favorite)
	favorite.ID = userinfo.Id
	code = model.SetUnImgFavorite(&favorite)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}
