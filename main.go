package main

import (
	"fmt"

	"github.com/Joyang0419/beego_note/routers"
	"github.com/astaxie/beego/config"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis" // Import Redis driver
	_ "github.com/go-sql-driver/mysql"                     // Import MySQL driver
)

func main() {
	var (
		conf config.Configer
		err  error
	)
	if conf, err = config.NewConfig("ini", "conf/app.conf"); err != nil {
		panic(fmt.Errorf("config.NewConfig err: %w", err))
	}

	MysqlConf := MYSQLConfig{
		Username: conf.String("dbuser"),
		Password: conf.String("dbpass"),
		Network:  conf.String("network"),
		IP:       conf.String("dbhost"),
		Port:     conf.String("dbport"),
		Dbname:   conf.String("dbname"),
	}

	if err = RegisterMYSQLORM(
		MysqlConf.Username,
		MysqlConf.Password,
		MysqlConf.Network,
		MysqlConf.IP,
		MysqlConf.Port,
		MysqlConf.Dbname,
		"default",
	); err != nil {
		panic(fmt.Errorf("RegisterMYSQLORM err: %w", err))
	}

	// init route
	routers.InitRoute()
	//

	web.BConfig.CopyRequestBody = true // 為了讓controller c.Ctx.Input.RequestBody 是有值的
	orm.Debug = true                   // 可能存在性能问题，不建议使用在生产模式, 可以印出sql語法
	web.Run()
}

type MYSQLConfig struct {
	Username string
	Password string
	Network  string
	IP       string
	Port     string
	Dbname   string
}

func RegisterMYSQLORM(username, password, network, ip, port, dbName, aliasName string) error {
	var (
		err        error
		dataSource = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, network, ip, port, dbName)
	)

	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		return fmt.Errorf("orm.RegisterDriver err: %w", err)
	}
	if err = orm.RegisterDataBase(aliasName, "mysql", dataSource); err != nil {
		return fmt.Errorf("orm.RegisterDataBase err: %w", err)
	}

	db, err := orm.GetDB(aliasName)
	if err != nil {
		return fmt.Errorf("orm.GetDB err: %w", err)
	}
	if err = db.Ping(); err != nil {
		return fmt.Errorf("db.Ping err: %w", err)
	}
	return nil
}
