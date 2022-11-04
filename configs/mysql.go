package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type MysqlConfig struct {
	Port   string `yaml:"port"`
	DBName string `yaml:"dbname"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Host   string `yaml:"host"`
}

func (m MysqlConfig) GetDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Pwd, m.Host, m.Port, "mysql",
	)
}
func InitMysql(configPath string) (db *gorm.DB, err error) {
	cnf, err := LoadFile(configPath)
	if err != nil {
		return nil, err
	}
	maps := cnf.Section("mysql")
	mysqlCnf := MysqlConfig{
		Port:   maps.Key("port").String(),
		DBName: maps.Key("dbname").String(),
		User:   maps.Key("user").String(),
		Pwd:    maps.Key("pwd").String(),
		Host:   maps.Key("host").String(),
	}
	db, err = gorm.Open(mysql.Open(mysqlCnf.GetDsn()))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func FirstInitMysql(configPath string) (db *gorm.DB, err error) {
	cnf, err := LoadFile(configPath)
	if err != nil {
		return nil, err
	}
	maps := cnf.Section("mysql")
	mysqlCnf := MysqlConfig{
		Port:   maps.Key("port").String(),
		DBName: maps.Key("initdb").String(),
		User:   maps.Key("user").String(),
		Pwd:    maps.Key("pwd").String(),
		Host:   maps.Key("host").String(),
	}
	db, err = gorm.Open(mysql.Open(mysqlCnf.GetDsn()))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitSmartBagMysql(configPath string) (db *gorm.DB, err error) {
	cnf, err := LoadFile(configPath)
	if err != nil {
		return nil, err
	}
	maps := cnf.Section("mysql")
	mysqlCnf := MysqlConfig{
		Port:   maps.Key("port").String(),
		DBName: maps.Key("db").String(),
		User:   maps.Key("user").String(),
		Pwd:    maps.Key("pwd").String(),
		Host:   maps.Key("host").String(),
	}
	db, err = gorm.Open(mysql.Open(mysqlCnf.GetDsn()))
	if err != nil {
		return nil, err
	}
	return db, nil
}
