package models

type Test struct {
	Id    uint64 `gorm: "PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Title string `gorm:"type:varchar(100);" json:"title"`
}

//type User struct {} // default table name is `users`
func (t Test) TableName() string {
	return "api_test"
}