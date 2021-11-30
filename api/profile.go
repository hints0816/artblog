package api

import (
	"hello/model"
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
			"status": code,
			"data":   data,
		},
	)
}
