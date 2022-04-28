package errcodes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// response 公共JSON响应结构体
type response struct {
	Code    ErrorInt               `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"` // 前端成功或者失败信息
	Error   string                 `json:"error"`   // 发生错误时，具体信息
}

// Response 实现 response
type Response interface {
	Success(*gin.Context)
	Failed(int, *gin.Context)
}

// New 创建一个 response 实例
func New(code ErrorInt, data map[string]interface{}, msg string, err string) Response {
	return &response{
		Code:    code,
		Data:    data,
		Message: msg,
		Error:   err,
	}
}

// Success 请求成功响应
func (r *response) Success(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	c.JSON(http.StatusOK, r)
}

// Failed 请求失败响应
func (r *response) Failed(status int, c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	c.JSON(status, r)
}
