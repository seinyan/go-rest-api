package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/seinyan/go-rest-api/configs"
	"time"
)

func NewDBConn(c configs.Database) (*gorm.DB, error) {
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DbName)

	var err error
	var conn *gorm.DB
	conn, err = gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	conn.DB().Ping()
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(100)
	conn.DB().SetConnMaxLifetime(time.Hour)

	return conn, nil
}
