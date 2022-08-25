package Model

import (
	uuid "github.com/satori/go.uuid"
	"go-study/Config"
	"go-study/Library/Gorm"
	"go-study/Library/MyTime"
	"gorm.io/gorm"
)

// 表名（未加表前缀）
const tableName = "user"

//
// User
// @Description: 表字段结构体
//
type User struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Uuid      uuid.UUID      `gorm:"column:uuid;not null;default:'';uniqueIndex:user_uuid_uindex;comment:'全局唯一标识'" json:"uuid"`
	Name      string         `gorm:"column:name;;not null;default:'';comment:'用户名'" json:"name"`
	Password  string         `gorm:"column:password;;not null;default:'';comment:'密码'" json:"password"`
	CreatedAt *MyTime.Time   `gorm:"column:created_at;index:user_created_at_index" json:"created_at"`
	UpdatedAt *MyTime.Time   `gorm:"column:updated_at;index:user_updated_at_index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"-"`
}

//
// TableName
// @Description: 表名 默认单表
// @receiver User
// @return string
//
func (User) TableName() string {
	prefix := Config.GetFieldByName(Config.Configs.Web, "DB", "Prefix")

	return prefix + tableName
}

// UserModel
// @Description: 初始化model 方便join查询
// @return *gorm.DB
//
func UserModel() *gorm.DB {
	return Gorm.Mysql.Table(User{}.TableName())
}
