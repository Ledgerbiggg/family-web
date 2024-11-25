package mysql

import (
	"context"
	"family-web-server/src/config"
	"family-web-server/src/log"
	"family-web-server/src/pkg/base"
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

func NewGorm(p base.Params) *GormDb {
	g := &GormDb{}
	g.config = p.Config
	g.log = p.Log
	address := p.Config.Mysql.Address
	username := p.Config.Mysql.Username
	password := p.Config.Mysql.Password
	database := p.Config.Mysql.Database
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, database)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		g.log.Error("failed to connect to database:" + err.Error())
		return nil
	}
	g.db = db
	return g
}
