package app

import (
	"gin/app/common/request"
	"gin/app/common/response"
	"gin/app/services"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var form request.Login
	if err := ctx.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(ctx, err.Error())
			return
		}
		response.Success(ctx, tokenData)
	}
}
