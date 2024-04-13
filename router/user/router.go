package user

import (
	"github.com/gin-gonic/gin"
	"kube-auto/api"
)

type LoginRouter struct {
}

func (*LoginRouter) Login(r *gin.Engine) {
	apiGroup := api.ApiGroupApp.LoginApiGroup
	r.POST("login", apiGroup.GetUserMessage)
}
