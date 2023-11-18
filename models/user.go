package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	ID        int       `orm:"column(id)"`
	Account   string    `orm:"column(account);size(255)"`
	Password  string    `orm:"column(password);size(255)"`
	LoginTime time.Time `orm:"column(login_time)"`
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
