package controller

import (
	"2022/short-url-service/model"
	"2022/short-url-service/service"
	"2022/short-url-service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VisitShortUrl(c *gin.Context) {
	surl := c.Param("surl")
	if len(surl) == 0 {
		c.Redirect(http.StatusFound, model.DefaultLink)
		return
	}

	lurl, err := service.UrlSvc.QueryLongLink(surl)
	if err != nil {
		c.Redirect(http.StatusFound, model.DefaultLink)
		return
	}
	// 302 跳转
	c.Redirect(http.StatusFound, lurl)
	return
}

func GenShortUrl(c *gin.Context) {
	longUrlParams := &model.LongUrlParams{}
	err := c.BindJSON(&longUrlParams)
	if err != nil {
		utils.ResponseError(c, 400, "post data error")
		return
	}
	lurl := longUrlParams.Url

	if len(lurl) == 0 || len(lurl) >= model.LUrlLongLimit {
		utils.ResponseError(c, 400, "lurl is empty or too long")
		return
	}
	surl, err := service.UrlSvc.GetShortUrl(lurl)
	if err != nil {
		utils.ResponseError(c, 500, err.Error())
		return
	}
	// 组装结果
	surl = fmt.Sprintf("%s%s%s", model.DomainNamePrefix, model.DomainNameMiddle, surl)
	utils.ResponseSuccess(c, surl)
}
