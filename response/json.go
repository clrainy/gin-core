package response

import (
	"github.com/clrainy/gin-core/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Json 响应操作
func Json(c *gin.Context, httpStatus int, code int, data interface{}) {
	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  constant.GetMsg(code),
		Data: data,
	})
}

// JsonSuccess 响应-成功
func JsonSuccess(c *gin.Context, data interface{}) {
	Json(c, http.StatusOK, constant.SUCCESS, data)
}

// JsonFail 响应-失败
func JsonFail(c *gin.Context, data interface{}) {
	Json(c, http.StatusOK, constant.ERROR, data)
}

// JsonPage 响应-分页
func JsonPage(c *gin.Context, httpCode int, code int, data interface{}, total, totalPage int) {
	c.JSON(httpCode, Paging{
		Code:      code,
		Msg:       constant.GetMsg(code),
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
}
