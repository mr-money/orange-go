package Database

import (
	"fmt"
	"github.com/shockerli/cvt"
	"go-study/Library/Gorm"
	"gorm.io/gorm"
	"log"
)

var migrations []map[string]interface{}

func init() {
	mysqlMigrations := getMysqlMigrations()
	migrations = append(migrations, mysqlMigrations...)

}

//
// InitMigrate
// @Description: AutoMigrate数据库自动迁移
//
func InitMigrate() {

	for _, migration := range migrations {
		migrateErr := dbSetTableOptions(
			cvt.String(migration["engine"]),
			cvt.String(migration["comment"]),
			cvt.String(migration["charset"]),
		).AutoMigrate(migration["model"])
		if migrateErr != nil {
			log.Println(migrateErr)
			return
		}
	}

	log.Println("Database " + Gorm.Mysql.Migrator().CurrentDatabase() + ":----- Migration Success!")

}

//
// dbSetTableOptions
// @Description: 表默认设置
// @param comment 表注释
// @param engine 表引擎
// @return *gorm.DB
//
func dbSetTableOptions(engine string, comment string, charset string) *gorm.DB {
	//设置表引擎和表注释
	setValue := fmt.Sprintf("ENGINE=%s COMMENT='%s' CHARSET='%s'", engine, comment, charset)

	//fmt.Println(setValue)

	return Gorm.Mysql.Set("gorm:table_options", setValue)
}
