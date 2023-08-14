package postgres

import (
	"fmt"

	"fbc-bookings/config"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pgOptions struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
}

func (p *pgOptions) getDns() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.dbName)
}

func NewConnection() *gorm.DB {
	dns := pgOptions{
		host:     config.Environments().Postgres.DbHost,
		port:     config.Environments().Postgres.DbPort,
		user:     config.Environments().Postgres.DbUser,
		password: config.Environments().Postgres.DbPassword,
		dbName:   config.Environments().Postgres.DbName,
	}

	dbInstance, err := gorm.Open(postgres.Open(dns.getDns()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	log.Info("Postgres Connection Successfully")
	return dbInstance
}
