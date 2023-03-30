package app

import (
	"gin/app/common/request"
	"gin/app/common/response"
	"gin/app/services"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var form request.Register
	if err := ctx.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		response.Success(ctx, user)
	}
}
