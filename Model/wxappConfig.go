package Model

import (
	"fmt"
	"gorm.io/gorm"
	"orange-go/Config"
	"orange-go/Library/Gorm"
	"orange-go/Library/MyTime"
)

// WxappConfig
// @Description: 微信小程序配置表
type WxappConfig struct {
	ID        int         `gorm:"primary_key" json:"id"` //
	Appid     string      `json:"appid"`                 //
	Appsecret string      `json:"appsecret"`             //
	CreatedAt MyTime.Time `json:"created_at"`            //
	UpdatedAt MyTime.Time `json:"updated_at"`            //
	Key       string      `json:"key"`                   //
	CertPem   string      `json:"cert_pem"`              //
	KeyPem    string      `json:"key_pem"`               //
}

// TableName
// @Description: 表名 默认单表
// @receiver WxappConfig
// @return string
func (WxappConfig) TableName() string {
	prefix := Config.GetFieldByName(Config.Configs.Web, "DB", "Prefix")

	return fmt.Sprintf("%s%s", prefix, "wxapp_config")
}

// Model
// @Description: 初始化model 方便join查询
// @return *gorm.DB
func (WxappConfig WxappConfig) Model() *gorm.DB {
	return Gorm.Mysql.Table(WxappConfig.TableName())
}

// GetOption
// @Description: 获取表基础配置
// @receiver WxappConfig
// @param key
// @return string
func (WxappConfig) GetOption(key string) string {
	option := map[string]string{
		"engine":  "InnoDB",
		"comment": "微信小程序配置表",
		"charset": "utf8mb4",
	}

	return option[key]
}
