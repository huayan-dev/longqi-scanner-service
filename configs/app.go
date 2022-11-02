package configs

import (
	"gopkg.in/ini.v1"
)

var APP *AppConfigs

type AppConfigs struct {
	Mode    string
	Host    string
	Port    string
	TcpHost string
	TcpPort string
}

func LoadFile(configPath string) (cnf *ini.File, err error) {
	cnf, err = ini.Load(configPath)
	if err != nil {
		return nil, err
	}
	return
}

func InitApp(configPath string) (cnf *AppConfigs, err error) {
	file, err := LoadFile(configPath)
	if err != nil {
		return &AppConfigs{}, err
	}
	cnf = &AppConfigs{
		Mode:    file.Section("app").Key("mode").String(),
		Host:    file.Section("app").Key("host").String(),
		Port:    file.Section("app").Key("port").String(),
		TcpHost: file.Section("app").Key("tcphost").String(),
		TcpPort: file.Section("app").Key("tcpport").String(),
	}
	return cnf, nil
}
