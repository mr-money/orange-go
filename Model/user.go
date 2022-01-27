package Model

import (
	"go-study/Library/Gorm"
	"time"
)

type User struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func init() {

	var userModel User
	Gorm.Connect().Model(&userModel)
}
