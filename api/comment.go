package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DelComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelComment(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func AddComment(c *gin.Context) {
	var comment model.Comment

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	_ = c.ShouldBindJSON(&comment)

	comment.UserId = userinfo.Id
	from_profile := model.GetProfileById(userinfo.Id)
	comment.Avatar = from_profile.Avatar
	if comment.CommentId != 0 {
		parentComment := model.GetParentComment(comment.CommentId)
		if parentComment.ParentId == 0 {
			comment.ParentId = comment.CommentId
		} else {
			comment.ParentId = parentComment.ParentId
		}
		comment.TouserId = parentComment.UserId
	}

	comment.Status = 1

	code = model.AddComment(&comment)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    comment,
		"message": errormsg.GetErrMsg(code),
	})
}

func GetCommentListFront(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	var usid uint

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	if userinfo == nil {
		usid = 0
	} else {
		usid = userinfo.Id
	}
	data, count, floorCount, code := model.GetCommentList(usid, id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"data":       data,
		"total":      count,
		"floorCount": floorCount,
		"message":    errormsg.GetErrMsg(code),
	})

}

func GetCommentListFrontAdmin(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	var usid int

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)

	if userinfo == nil {
		usid = -1
	} else {
		usid = int(userinfo.Id)
	}

	data, count, code := model.GetCommentListAdmin(usid, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   count,
		"message": errormsg.GetErrMsg(code),
	})

}

func SetCommentDigg(c *gin.Context) {
	var digg model.Digg
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&digg)
	digg.ID = userinfo.Id
	code = model.SetDigg(&digg)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func SetCommentUnDigg(c *gin.Context) {
	var digg model.Digg
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&digg)
	digg.ID = userinfo.Id
	code = model.SetUnDigg(&digg)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}
