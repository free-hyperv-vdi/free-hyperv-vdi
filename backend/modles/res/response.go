package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Err     = 2
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}


func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "查询成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}


func Fail(data any, msg string, c *gin.Context) {
	Result(Err, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Err, map[string]any{}, msg, c)
}
