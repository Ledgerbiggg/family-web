package mysql

import (
	"context"
	"family-web-server/src/config"
	l "family-web-server/src/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type GormDb struct {
	db     *gorm.DB
	config *config.GConfig
	log    *l.ConsoleLogger
	ctx    context.Context
}

func (g *GormDb) GetDb() *gorm.DB {
	return g.db
}

func NewGorm(
	l *l.ConsoleLogger,
	c *config.GConfig,
) *GormDb {
	g := &GormDb{}
	g.config = c
	g.log = l
	address := c.Mysql.Address
	username := c.Mysql.Username
	password := c.Mysql.Password
	database := c.Mysql.Database
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, database)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		// 设置日志级别为 Info，显示 SQL 语句
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold: time.Second, // 慢查询阈值
				LogLevel:      logger.Info, // 记录所有SQL查询
				Colorful:      true,        // 启用彩色输出
			},
		),
	})
	if err != nil {
		g.log.Error("failed to connect to database:" + err.Error())
		return nil
	}
	g.db = db
	return g
}
