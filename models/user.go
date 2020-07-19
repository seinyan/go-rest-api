package models

type User struct {
	Id           uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	Name         string `gorm:"type:varchar(100);" json:"name" binding:"required" `
	Email        string `gorm:"type:varchar(100);NOT NULL;" json:"email" binding:"required,email"`
	PasswordHash string `gorm:"type:varchar(255);NOT NULL;" json:"password_hash"`
	Password     string `gorm:"-" json:"password" binding:"required,min=6,max=64"`
}

//type User struct {} // default table name is `users`
func (u User) TableName() string {
	return "api_user"
}

func (u *User) Sanitize() {

}