package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Query("id"))
	userId, _ := strconv.Atoi(c.Query("user_id"))
	keyword := c.Query("keyword")

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	edit := false
	if userinfo != nil {
		if userId == int(userinfo.Id) {
			edit = true
		}
	}

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 15
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.ListArticle(edit, userId, cid, keyword, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"edit":    edit,
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
	// file, fileHeader, _ := c.Request.FormFile("file")
	// fileSize := fileHeader.Size
	// contentType := fileHeader.Header.Get("Content-Type")
	// model.UpLoadFile(file, contentType, fileSize)
	// err := c.Bind(&data)
	// fmt.Println(err)
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	_ = c.ShouldBindJSON(&data)

	src := data.Desc
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	if len(src) < 200 {
		data.Desc = src
	} else {
		data.Desc = data.Content[0:200]
	}
	data.UserID = userinfo.Id
	if data.ID == 0 {
		code = model.CreateArt(&data)
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
	profile := model.GetProfileById(data.UserID)
	data.Profile = profile
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
