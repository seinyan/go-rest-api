package models

const (
	UserRoleUser    = 1
	UserRoleAdmin   = 2
	UserRoleManager = 3
)

type UserRegister struct {
	Username string `binding:"required,email,min=6,max=64" json:"username"`
	Password string `binding:"required,min=6,max=64" json:"password"`
}

type UserPerson struct {
	BaseModel
	FirstName  *string `gorm:"type:varchar(100);" json:"name" binding:"required"`
	LastName   *string `gorm:"type:varchar(100);" json:"name" binding:"required"`
	MiddleName *string `gorm:"type:varchar(100);" json:"name" binding:"required"`
}

type User struct {
	BaseModel
	Phone        *string `gorm:"type:varchar(100);" json:"phone"`
	Email        string  `gorm:"type:varchar(100);NOT NULL;" json:"email" binding:"required,email"`
	PasswordHash string  `gorm:"type:varchar(255);NOT NULL;" json:"-"`
	Password     string  `gorm:"-" json:"password" binding:"required,min=6,max=64"`
	Role         string  `gorm:"type:varchar(100);" json:"role"`
	IsActive     bool    `gorm:"type:boolean;DEFAULT:true"`
	Person       UserPerson  `gorm:"ForeignKey:PersonID" json:"person"`
	PersonID     uint64  `json:"-"`
}

//type User struct {} // default table name is `users`
func (u User) TableName() string {
	return "api_user"
}
func (p UserPerson) TableName() string {
	return "api_user_person"
}

func (u *User) Sanitize() {}