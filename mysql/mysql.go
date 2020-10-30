// @Description  mysql
// @Author  	 jiangyang  
// @Created  	 2020/10/30 3:44 下午
package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// mysql连接
	db *gorm.DB
	// 保证只建立一次连接
	once sync.Once
)

// Mysql配置结构体
type Config struct {
	User        string `json:"user" yaml:"user"`                   // 用户名
	Password    string `json:"password" yaml:"password"`           // 密码
	Host        string `json:"host" yaml:"host"`                   // 主机地址
	Port        int    `json:"port" yaml:"port"`                   // 端口号
	Dbname      string `json:"dbname" yaml:"dbname"`               // 数据库名
	MaxIdleConn int    `json:"max_idle_conn" yaml:"max_idle_conn"` // 最大空闲连接
	MaxOpenConn int    `json:"max_open_conn" yaml:"max_open_conn"` // 最大活跃连接
	Debug       bool   `json:"debug" yaml:"debug"`                 // 是否开启Debug（开启Debug会打印数据库操作日志）
}

// 初始化数据库
func Init(mysqlConfig Config) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			mysqlConfig.User,
			mysqlConfig.Password,
			mysqlConfig.Host,
			mysqlConfig.Port,
			mysqlConfig.Dbname,
		)

		conn, err := gorm.Open("mysql", dsn)
		if err != nil {
			logrus.Fatalf("mysql connect failed: %v", err)
		}

		conn.DB().SetMaxIdleConns(viper.GetInt("db.mysql.max_idle_conn"))
		conn.DB().SetMaxOpenConns(viper.GetInt("db.mysql.max_open_conn"))

		conn.LogMode(viper.GetBool("db.mysql.debug"))

		if err = conn.DB().Ping(); err != nil {
			logrus.Fatalf("database heartbeat failed: %v", err)
		}
		db = conn
		logrus.Info("mysql connect successfully")
	})
}

// 获取Mysql连接
func Conn() *gorm.DB {
	return db
}

// Close method
func Close() error {
	if db != nil {
		if err := db.Close(); err != nil {
			return errors.WithStack(err)
		}
	}

	logrus.Info("mysql connect closed")
	return nil
}
