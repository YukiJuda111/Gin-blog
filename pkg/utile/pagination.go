package util

import (
	"strconv"

	"github.com/YukiJuda111/go-gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

// GetPage get index of the current page
func GetPage(c *gin.Context) int {
	result := 0
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
