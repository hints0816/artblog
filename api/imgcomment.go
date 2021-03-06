package api

import (
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddImgComment(c *gin.Context) {
	var comment model.Imgcomment
	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&comment)
	comment.UserId = userinfo.Id
	from_profile := model.GetProfileById(userinfo.Id)
	comment.Avatar = from_profile.Avatar
	if comment.CommentId != 0 {
		parentComment := model.GetParentImgComment(comment.CommentId)
		if parentComment.ParentId == 0 {
			comment.ParentId = comment.CommentId
		} else {
			comment.ParentId = parentComment.ParentId
		}
		comment.TouserId = parentComment.UserId
	}
	comment.Status = 1
	code = model.AddImgComment(&comment)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    comment,
		"message": errormsg.GetErrMsg(code),
	})
}

func GetImgCommentList(c *gin.Context) {
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
	data, count, floorCount, code := model.GetImgCommentList(usid, id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":     code,
		"data":       data,
		"total":      count,
		"floorCount": floorCount,
		"message":    errormsg.GetErrMsg(code),
	})

}
