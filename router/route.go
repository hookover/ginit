package router

import (
	"gin_api/app/http/controllers"
	"gin_api/app/http/controllers/jaeger_conn"
	"gin_api/app/http/middlewares/exception"
	"gin_api/app/http/middlewares/jaeger"
	"gin_api/app/http/middlewares/logger"
	"gin_api/app/http/middlewares/requestid"
	"gin_api/util/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), jaeger.SetUp(), requestid.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	engine.GET("/api/v1/auth/login", controllers.Login)
	engine.GET("/api/v1/auth/register", controllers.Register)

	// 测试链路追踪
	engine.GET("/jaeger_test", jaeger_conn.JaegerTest)
}
