package api

import (
	"fmt"
	"hello/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": 1,
		},
	)
}

// GitHub OAuth App Callback
func GithubCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println(code)
	fmt.Println(state)
	c.JSON(
		http.StatusOK, gin.H{
			"code":  code,
			"state": state,
		},
	)

}
