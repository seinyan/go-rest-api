package migrations

import (
	"fmt"
)

func Migrate() {
	fmt.Println("Migrate ... NOT Working")
	//database.DBConn.AutoMigrate(models.User{}, models.Test{})
	fmt.Println("Migrate end ...")
}