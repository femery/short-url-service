package component

import (
	"2022/short-url-service/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql" //引入mysql驱动
)

var (
	once sync.Once
	DB   *gorm.DB
)

// InitDB 初始化DB
func initDB(dbConfig *conf.MySQLConf) error {

	var err error
	// 全局只执行一次
	once.Do(func() {
		connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?"+
			"charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s", dbConfig.User,
			dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)

		fmt.Println(connArgs)

		DB, err = gorm.Open("mysql", connArgs)
		if err != nil {
			return
		}

		// SetMaxIdleConns 设置空闲连接池中的最大连接数。
		DB.DB().SetMaxIdleConns(10)

		// SetMaxOpenConns 设置数据库连接最大打开数。
		DB.DB().SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置可重用连接的最长时间
		DB.DB().SetConnMaxLifetime(time.Hour)

		DB.SingularTable(true)

		DB.LogMode(true)

	})
	return err
}

func InitDBByCfg(cfg *conf.Config) (err error) {
	return initDB(cfg.MySQLConf)
}
