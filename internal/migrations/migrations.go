package migrations

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/seinyan/go-rest-api/internal/models"
)

func Migrate(db *gorm.DB) {
	fmt.Println("Migrate ...")

	db.AutoMigrate(models.Person{}, models.User{})
	db.Model(&models.User{}).AddForeignKey("person_id", "api_person(id)", "CASCADE", "CASCADE")
	db.Model(&models.User{}).AddUniqueIndex("api_person_id_unique", "person_id")


	fmt.Println("Migrate end ...")
}