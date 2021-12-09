package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Query("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 15
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.ListArticle(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

func ListArticleRepository(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Query("id"))

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 15
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.ListArticleRepository(userinfo.Id, id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	var data1 model.ArticleVo
	// file, fileHeader, _ := c.Request.FormFile("file")
	// fileSize := fileHeader.Size
	// contentType := fileHeader.Header.Get("Content-Type")
	// model.UpLoadFile(file, contentType, fileSize)
	// err := c.Bind(&data)
	// fmt.Println(err)
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	_ = c.ShouldBindJSON(&data)
	_ = c.ShouldBindJSON(&data1)
	data.Desc = data.Content[0:19]
	data.UserID = userinfo.Id
	if data.ID == 0 {
		code = model.CreateArt(&data1)
	} else {
		code = model.UpdateArt(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

// DelArticle 删除文章
func DelArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelArt(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

// GetUsers 查询用户列表
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

func UploadArtImage(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	contentType := fileHeader.Header.Get("Content-Type")
	url, code := model.UpLoadFile(file, contentType, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"url":     url,
		"message": errormsg.GetErrMsg(code),
	})
}
