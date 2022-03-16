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
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"column:name;" json:"name"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"-"`
}

var UserModel *gorm.DB

func init() {
	var user User
	UserModel = Gorm.Mysql.Table(tableName).Model(&user)

}
