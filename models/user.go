package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id       int
	Account  string
	Password string
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(User),
	)
}
