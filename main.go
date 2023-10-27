package main

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/beego/beego/v2/core/config/yaml" // Import yaml
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func main() {
	var (
		err       error
		conf      config.Configer
		MysqlConf MYSQLConfig
	)

	if conf, err = config.NewConfig("yaml", "conf/app.yaml"); err != nil {
		panic(fmt.Errorf("config.NewConfig err: %w", err))
	}

	if err = conf.Unmarshaler("mysql", &MysqlConf); err != nil {
		panic(fmt.Errorf("conf.Unmarshaler err: %w", err))
	}

	o, err := RegisterMYSQLORM(
		MysqlConf.Username,
		MysqlConf.Password,
		MysqlConf.Network,
		MysqlConf.IP,
		MysqlConf.Port,
		MysqlConf.DbName,
		"default",
	)
	if err != nil {
		panic(fmt.Errorf("RegisterMYSQLORM err: %w", err))
	}

	_ = o
	web.Run()
}

type MYSQLConfig struct {
	Username string
	Password string
	Network  string
	IP       string
	Port     string
	DbName   string
}

func RegisterMYSQLORM(username, password, network, ip, port, dbName, aliasName string) (orm.Ormer, error) {
	fmt.Println(port)
	var (
		err        error
		dataSource = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, network, ip, port, dbName)
	)

	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		return nil, fmt.Errorf("orm.RegisterDriver err: %w", err)
	}
	if err = orm.RegisterDataBase(aliasName, "mysql", dataSource); err != nil {
		return nil, fmt.Errorf("orm.RegisterDataBase err: %w", err)
	}

	db, err := orm.GetDB(aliasName)
	if err != nil {
		return nil, fmt.Errorf("orm.GetDB err: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping err: %w", err)
	}

	return orm.NewOrmWithDB("mysql", aliasName, db)
}
