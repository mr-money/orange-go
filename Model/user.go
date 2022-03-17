package Model

import (
	uuid "github.com/satori/go.uuid"
	"go-study/Library"
	"go-study/Library/Gorm"
	"gorm.io/gorm"
)

//表名
var tableName = "user"

//
// User
// @Description: 表字段结构体
//
type User struct {
	ID        uint64        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Uuid      uuid.UUID     `gorm:"column:uuid;" json:"uuid"`
	Name      string        `gorm:"column:name;" json:"name"`
	CreatedAt *Library.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt *Library.Time `gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt *Library.Time `gorm:"column:deleted_at;" json:"-"`
}

// UserModel
// @Description: 初始化model
// @return *gorm.DB
//
func UserModel() *gorm.DB {
	return Gorm.Mysql.Table(tableName)
}
