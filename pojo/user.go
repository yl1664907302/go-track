package pojo

type User struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	RoleId      string `json:"roleId"`
	Permissions string `json:"permissions"`
}
