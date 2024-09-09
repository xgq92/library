package data

import (
	"fmt"
	"os"
	"sync"
	"time"

	logger "github.com/alecthomas/log4go"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**********************************
1.初始化数据
***********************************/

// 数据初始化对象
var DataInitIns = new(DataInit)

type DataInit struct {
	SqlDb       *gorm.DB
	mySQLConfig MySQLConfig
	redisDB     *redis.Client
	redisOnce   sync.Once
	sync.Mutex
}

type MySQLConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	Database string
	PoolSize int
}

func (this *DataInit) OnInit() {
	defer logger.Info("DataInit OnInit ok")
	// mysql数据库init
	this.mySQLConfig = MySQLConfig{}
	err := this.connectSqlDb()
	if err != nil {
		logger.Error(`DataInit::init connectSqlDb error: %v`, err)
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}
	// redis
	this.setupRedis()

	//other init
}

func (this *DataInit) OnClose() {
	// 进行其它数据相关去初始化
	if this.SqlDb != nil {
		this.SqlDb.Close()
		logger.Info("DataInit::OnClose Close SqlDb")
	}
	if this.redisDB != nil {
		this.redisDB.Close()
		logger.Info("DataInit::OnClose Close redisDB")
	}
}

func (this *DataInit) GetSqlDb() *gorm.DB {
	this.Lock()
	defer this.Unlock()
	if this.SqlDb == nil {
		this.connectSqlDb()
		return this.SqlDb
	}
	err := this.SqlDb.DB().Ping()
	if err != nil {
		logger.Warn("DataInit::connectSqlDb 数据库连接失败重试:", err)
		logger.Warn("DataInit::connectSqlDb 数据库连接重试!")
		this.SqlDb.Close()
		this.connectSqlDb()
		return this.SqlDb
	}
	return this.SqlDb
}

func (this *DataInit) connectSqlDb() error {
	var err error
	this.SqlDb, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Error("DataInit::connectSqlDb 数据库连接失败:", err)
		return err
	}
	this.SqlDb.DB().SetMaxIdleConns(5)
	this.SqlDb.DB().SetMaxOpenConns(10)
	if err := this.SqlDb.DB().Ping(); err != nil {
		fmt.Println("DataInit::connectSqlDb ping数据库连接失败:", err)
		return err
	}

	// 自动迁移
	// db.AutoMigrate(&UserInfo{})

	return nil
}

func (this *DataInit) setupRedis() {
	this.redisOnce.Do(func() {
		this.redisDB = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379", // Redis地址
			Password: "",               // Redis密码，如果没有则为空字符串
			DB:       0,                // 使用默认DB
		})
	})
}

func GetDb() *gorm.DB {
	return DataInitIns.GetSqlDb()
}

func GetRedis() *redis.Client {
	return DataInitIns.redisDB
}
