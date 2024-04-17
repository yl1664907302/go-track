package user

import (
	"github.com/gin-gonic/gin"
	"kube-auto/database/service"
	"kube-auto/pojo"
	"kube-auto/response"
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

	// 表单数据已成功绑定到loginForm，现在可以访问和使用这些字段
	username := loginForm.Username
	password := loginForm.Password

	user, err := service.UserServiceImpl.GetUserByNameAndPasswd(username, password)
	if err != nil {
		log.Print(err)
	}

	if username == "" && password == "" {
		response.FailWithDetailed(c, "用户登入失败", map[string]string{
			"message": "用户：" + username + " 登入失败！",
		})

	} else if user.Username == username && user.Password == password {
		response.LoginSuccessDetailed(c, "用户："+username+" 登入成功！", map[string]string{
			"token": "123456",
		})

	} else {
		response.FailWithDetailed(c, "用户登入失败", map[string]string{
			"message": "用户：" + username + " 登入失败！",
		})
	}
}
