package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/seinyan/go-rest-api/configs"
	"github.com/seinyan/go-rest-api/internal/database"
	"github.com/seinyan/go-rest-api/internal/migrations"
	"log"
	"time"
)

func main() {
	c, err:= configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	db, err := database.NewDBConn(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	migrations.Migrate(db)

	time.Sleep(time.Second * 5)
}