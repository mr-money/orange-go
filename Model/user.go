package Model

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"orange-go/Config"
	"orange-go/Library/Gorm"
	"orange-go/Library/MyTime"
)

// User
// @Description: 表字段结构体
type User struct {
	ID        uint64         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Uuid      uuid.UUID      `gorm:"column:uuid;type:varchar(50);not null;default:'';uniqueIndex:user_uuid_index;comment:全局唯一标识" json:"uuid"`
	Name      string         `gorm:"column:name;type:varchar(20);not null;default:'';comment:用户名" json:"name"`
	Password  string         `gorm:"column:password;type:varchar(64);not null;default:'';comment:密码" json:"password"`
	CreatedAt *MyTime.Time   `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;index:user_created_at_index" json:"created_at"`
	UpdatedAt *MyTime.Time   `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;index:user_updated_at_index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"-"`
}

// TableName
// @Description: 表名 默认单表
// @receiver User
// @return string
func (User) TableName() string {
	prefix := Config.GetFieldByName(Config.Configs.Web, "DB", "Prefix")

	return fmt.Sprintf("%s%s", prefix, "user")
}

// Model
// @Description: 初始化model 方便join查询
// @return *gorm.DB
func (user User) Model() *gorm.DB {
	return Gorm.Mysql.Table(user.TableName())
}

// GetOption
// @Description: 获取表基础配置
// @receiver User
// @param key
// @return string
func (User) GetOption(key string) string {
	option := map[string]string{
		"engine":  "InnoDB",
		"comment": "用户表",
		"charset": "utf8mb4",
	}

	return option[key]
}
