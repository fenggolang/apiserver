package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// 因为一个API服务器可能需要同时访问多个数据库，为了对多个数据库进行初始化
// 和连接管理，这里定义了一个叫Database的struct
type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

// 单例设计模式，下面DB即为单例对象
var DB *Database

// openDB方法建立数据库连接，不同数据库实例传入不同的username、
// password、addr和名字信息，从而建立不同的数据库连接
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	// openDB()最终调用gorm.Open()来建立一个数据库连接
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(viper.GetInt("db.max_open_conns"))
	db.DB().SetMaxIdleConns(viper.GetInt("db.max_idle_conns"))
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

// Database结构体的Init()方法用来初始化连接
// Init()方法会调用GetSelfDB()和GetDockerDB()函数来同时创建两个Database的数据库对象
func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
		Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}
