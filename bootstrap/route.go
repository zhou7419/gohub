package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	// 全局中间件
	registerGlobalMiddleWare(router)

	// 路由
	routes.ResgsterAPIRoutes(router)

	// 处理 404
	setup404Handle(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handle(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 html
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 json
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义",
			})
		}
	})
}
