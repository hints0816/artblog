package api

import (
	"hello/model"
	"hello/utils/errormsg"
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

	_ = c.ShouldBindJSON(&data)
	code = model.PostArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}
