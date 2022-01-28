package Model

import (
	"go-study/Library/Gorm"
	"gorm.io/gorm"
	"time"
)

var tableName = "user"

type User struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

var UserModel *gorm.DB

func init() {
	var user User
	UserModel = Gorm.DBInstance.Table(tableName).Model(&user)
}
