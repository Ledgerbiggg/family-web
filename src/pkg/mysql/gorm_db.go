package mysql

import (
	"context"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDb struct {
	db     *gorm.DB
	config *config.GConfig
	log    *log.ConsoleLogger
	ctx    context.Context
}

func (g *GormDb) GetDb() *gorm.DB {
	return g.db
}

func NewGorm(
	l *log.ConsoleLogger,
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
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		g.log.Error("failed to connect to database:" + err.Error())
		return nil
	}
	g.db = db
	return g
}
