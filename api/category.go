package api

import (
	"hello/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
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

	data, total := model.GetCate(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"data":  data,
		"total": total,
	})
}
