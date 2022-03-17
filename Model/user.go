package Model

import (
	uuid "github.com/satori/go.uuid"
	"go-study/Library/Gorm"
	"go-study/Library/MyTime"
	"gorm.io/gorm"
)

//表名
var tableName = "user"

//
// User
// @Description: 表字段结构体
//
type User struct {
	ID        uint64       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Uuid      uuid.UUID    `gorm:"column:uuid;uniqueIndex:user_uuid_uindex" json:"uuid"`
	Name      string       `gorm:"column:name;" json:"name"`
	CreatedAt *MyTime.Time `gorm:"column:created_at;index:user_created_at_index" json:"created_at"`
	UpdatedAt *MyTime.Time `gorm:"column:updated_at;index:user_updated_at_index" json:"updated_at"`
	DeletedAt *MyTime.Time `gorm:"column:deleted_at;" json:"-"`
}

// UserModel
// @Description: 初始化model
// @return *gorm.DB
//
func UserModel() *gorm.DB {
	return Gorm.Mysql.Table(tableName)
}
