package pojo

type User struct {
	Username    string `gorm:"column:username"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	Role        string `gorm:"column:role"`
	RoleId      string `gorm:"column:roleId"`
	Permissions string `gorm:"column:permissions"`
}
