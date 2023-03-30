package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
	TokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{4000, "业务错误"},
	ValidateError: CustomError{4100, "请求参数错误"},
	TokenError:    CustomError{4200, "登录授权失效"},
}
