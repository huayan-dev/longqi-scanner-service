package main

import (
	"flag"
	"fmt"
	"log"
	"scanner-service/configs"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "../configs/config.ini", "config file path")
	flag.Parse()
	fileName := fmt.Sprintf("%v-%v-%v.log",
		time.Now().Format("2006"),
		time.Now().Format("01"),
		time.Now().Format("02"),
	)
	log.SetOutput(initLog(fileName))
	initConfigs()
	initDatabase()
	deleteLog()
}
func main() {
	r := loadRouters()
	address := fmt.Sprintf("%s:%s", configs.APP.Host, configs.APP.Port)
	err := r.Run(address)
	if err != nil {
		panic(err)
	}
}
