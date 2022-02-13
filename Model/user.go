package Model

import (
	"go-study/Library/Gorm"
	"gorm.io/gorm"
	"time"
)

//表名
var tableName = "user"

//
// User
// @Description: 表字段结构体
//
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
	UserModel = Gorm.Mysql.Table(tableName).Model(&user)
}
