package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"io"
	"io/fs"
	"log"
	"os"
	"scanner-service/configs"
	"scanner-service/models"
	"strconv"
	"strings"
	"time"
)

/**
  @author:pandi
  @date:2022-11-03
  @note:
**/
func migrate(tx *gorm.DB) (err error) {
	err = tx.AutoMigrate(&models.MaterialPackages{})
	return err
}
func initLog(logFileName string) (w io.Writer) {
	logPath := "./logs/"
	if !dirExists(logPath) {
		os.Mkdir(logPath, os.ModePerm)
	}
	fullPath := logPath + logFileName
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		f, err := os.Create(fullPath)
		if err != nil {
			panic("log file create failed:" + fullPath)
		}
		f.Close()
	}
	logFile, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(logFile)
	return logFile
}
func dirExists(path string) bool {
	i, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if i.IsDir() && err == nil {
		//删除空目录. os.remove对于目录如果非空则不执行，只删除空目录
		//err := os.Remove(path)
		//if err != nil {
		//	log.Println("dirExists remove", err)
		//}
		return true
	}
	return false
}
func initDatabase() {
	_ = configs.DB.Transaction(func(tx *gorm.DB) error {
		log.Println("1.Start init database...")
		var dbname []string
		err := tx.Table("innodb_table_stats").Select("database_name").Where("database_name = ?", "longqi").Scan(&dbname).Error
		//err := tx.Exec("SHOW DATABASES LIKE 'longqi';").Error
		if err != nil {
			log.Println("SHOW DATABASES LIKE 'longqi' %v:" + err.Error())
			return err
		}
		//log.Println("dbname = " + dbname)
		if len(dbname) == 0 {
			//初始化数据库、表结构
			err := tx.Exec("create database longqi;").Error
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			log.Println("Database create success!")
			log.Println("2. Start init tables...")
			tx.Exec("use longqi;")
			err = migrate(tx)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			log.Println("Tables create success!")
		}
		tx.Exec("use longqi;")
		//新建默认用户
		//fmt.Println("3. Start init default user...")
		//user, password, err := createAdminUser(tx)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return err
		//}
		//fmt.Printf("Default user created! account:%s password:%s \n", user.Account, password)
		return nil
	})
}
func initConfigs() {
	db, err := configs.InitMysql(configPath)
	if err != nil {
		panic(err)
	}
	configs.DB = db

	appCnf, err := configs.InitApp(configPath)
	if err != nil {
		panic(err)
	}
	configs.APP = appCnf
}
func deleteLog() {
	//添加定时任务
	crontab := cron.New()
	task := func() {
		deleteLogDays()
	}
	_, err := crontab.AddFunc("@daily", task)
	if err != nil {
		log.Fatal(err)
	}
	crontab.Start()
	defer crontab.Stop()
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {

	}
}
func deleteLogDays() {
	days, _ := strconv.Atoi(configs.APP.DeleteLogDays)
	root := os.DirFS("./logs")
	err := fs.WalkDir(root, ".", func(backPath string, d fs.DirEntry, err error) error {
		logName := d.Name()
		if logName == "." {
			return nil
		}
		names := strings.Split(logName, ".")
		//os.Setenv()
		local, err := time.LoadLocation("Asia/Beijing")
		if err != nil {
			local = time.FixedZone("CST", 8*3600)
		}
		logTime, _ := time.ParseInLocation("2006-01-02", names[0], local)
		//log.Printf("logTime : %v\n", logTime)
		if time.Now().Unix()-logTime.Unix() > int64(days*24*3600) {
			err := os.Remove("./logs/" + logName)
			if err != nil {
				log.Printf("删除日志:%s\n失败", logName)
				return err
			} else {
				log.Printf("删除日志:%s\n成功", logName)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("deleteLogDays() fs.WalkDir err :%v\n", err)
	}
}
