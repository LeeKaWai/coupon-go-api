package api

import (
	"coupon-go-api/src/service"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type QueryModel struct {
	KeyWord string `json:"q"`
}

// 淘宝搜索
func TaobaoSearch(c *gin.Context) {
	json := QueryModel{}
	c.BindJSON(&json)
	res, err := service.TKLSearch(url.QueryEscape(json.KeyWord), "")
	if err != nil {
		fmt.Printf("%s", err)
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
