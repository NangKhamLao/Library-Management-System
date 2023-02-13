package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type mysqlDB struct {
	Host string
	Port string
	User string
	Pass string
	DB   string
}

var (
	MysqlDB mysqlDB
)

func init() {
	if os.Getenv("TESTING") == "1" {
		return
	}
	iniPath := "config/config.ini"
	if args := os.Args; len(args) > 1 {
		iniPath = args[1]
	}

	iniFile, err := ini.Load(iniPath)
	if err != nil {
		log.Fatalf("load %s error:%s\n", iniPath, err.Error())
		os.Exit(1)
	}

	//mysql
	database := iniFile.Section("mysql")
	MysqlDB.Host = database.Key("MysqlHost").String()
	MysqlDB.Port = database.Key("MysqlPort").String()
	MysqlDB.User = database.Key("MysqlUser").String()
	MysqlDB.Pass = database.Key("MysqlPass").String()
	MysqlDB.DB = database.Key("MysqlDB").String()

}
