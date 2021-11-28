package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	var data model.Comment

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	_ = c.ShouldBindJSON(&data)

	data.Username = userinfo.Username
	data.UserId = userinfo.Id
	data.Status = 1

	code = model.AddComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

func AddCommentChild(c *gin.Context) {
	var data model.Commentchild
	var commentchildVo model.CommentchildVo

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	_ = c.ShouldBindJSON(&commentchildVo)
	_ = c.ShouldBindJSON(&data)

	data.From = userinfo.Id

	code = model.AddCommentChild(&data, &commentchildVo)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}

func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetCommentListFront(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})

}
