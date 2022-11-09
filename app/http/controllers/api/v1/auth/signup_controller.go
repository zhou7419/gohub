package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExisRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExists(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
