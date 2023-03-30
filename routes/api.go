package routes

import (
	"gin/app/controllers/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(route *gin.RouterGroup) {
	route.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	route.GET("/test", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "success")
	})

	route.POST("/auth/register", app.Register)
	route.POST("/auth/login", app.Login)

	//route.POST("/user/register", func(context *gin.Context) {
	//	var form request.Register
	//	if err := context.ShouldBindJSON(&form); err != nil {
	//		context.JSON(http.StatusOK, gin.H{
	//			"error": request.GetErrorMsg(form, err),
	//		})
	//		return
	//	}
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "success",
	//	})
	//})
}
