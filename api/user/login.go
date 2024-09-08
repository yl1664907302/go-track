package user

import (
	"github.com/gin-gonic/gin"
	"go-track/database/mysql"
	"go-track/pojo"
	"go-track/response"
	"log"
	"net/http"
)

type LoginApi struct {
}

func (*LoginApi) GetUserMessage(c *gin.Context) {
	var loginForm pojo.User
	// 使用ShouldBind方法将请求上下文（c）中的表单数据绑定到LoginForm实例
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	user, err := mysql.LoginUser(loginForm.Username, loginForm.Password)
	if err != nil {
		log.Print(err)
		response.FailWithDetailed(c, "用户登入失败", map[string]any{
			"code": http.StatusInternalServerError,
		})

	} else {
		response.LoginSuccessDetailed(c, "登入成功！", map[string]any{
			"username":    user.Username,
			"role":        user.Role,
			"roleId":      user.RoleId,
			"permissions": user.Permissions,
		})

	}
}
