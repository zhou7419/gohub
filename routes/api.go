package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResgsterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)

			authGroup.POST("/signup/phone/exists", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exists", suc.IsEmailExists)
		}
	}
}
