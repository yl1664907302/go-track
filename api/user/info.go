package user

import (
	"github.com/gin-gonic/gin"
	"kube-auto/response"
)

type InfoApi struct {
}

func (*InfoApi) GetUserInfo(c *gin.Context) {

	response.InfoSuccessDetailed(c, "登入成功", map[string]string{
		"roles":  "admin",
		"name":   "admin",
		"avatar": "xxxx",
	})
}
