package user

import (
	"github.com/gin-gonic/gin"
	"kube-auto/response"
)

type InfoApi struct {
}

func (*InfoApi) GetUserInfo(c *gin.Context) {
	response.SuccssWithDetailed(c, "用户登入失败", map[string]string{
		"message": " 登入失败！",
	})
}
