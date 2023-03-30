package response

import (
	"gin/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构体
type Response struct {
	ErrorCode int         `json:"error_code"` // 自定义错误码
	Data      interface{} `json:"data"`       // 数据
	Message   string      `json:"message"`    // 信息
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{0, data, "ok"})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(ctx *gin.Context, errorCode int, msg string) {
	ctx.JSON(http.StatusOK, Response{errorCode, nil, msg})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(ctx *gin.Context, customError global.CustomError) {
	Fail(ctx, customError.ErrorCode, customError.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(ctx *gin.Context, msg string) {
	Fail(ctx, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(ctx *gin.Context, msg string) {
	Fail(ctx, global.Errors.BusinessError.ErrorCode, msg)
}

// TokenFail 鉴权失败
func TokenFail(ctx *gin.Context) {
	FailByError(ctx, global.Errors.TokenError)
}
