package pojo

type User struct {
	Username string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"` // 实际应用中应存储加密后的密码
	Role     string // 可以定义枚举类型或外键关联角色表
}
