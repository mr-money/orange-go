package Database

import (
	"fmt"
	"github.com/shockerli/cvt"
	"gorm.io/gorm"
	"orange-go/Library/Gorm"
	"orange-go/Library/Logger"
)

var dbLogger = Logger.MustModuleLogger("database")

// InitMigrate
// @Description: AutoMigrate数据库自动迁移
func InitMigrate() {
	mysqlMigrations := getMysqlMigrations()

	var migrations []map[string]interface{}
	migrations = append(migrations, mysqlMigrations...)

	for _, migration := range migrations {
		migrateErr := dbSetTableOptions(
			cvt.String(migration["engine"]),
			cvt.String(migration["comment"]),
			cvt.String(migration["charset"]),
		).AutoMigrate(migration["model"])
		if migrateErr != nil {
			dbLogger.Error("migration error", "error", migrateErr)
			return
		}
	}

	dbLogger.Info("Database migration success", "database", Gorm.Mysql.Migrator().CurrentDatabase())

}

// dbSetTableOptions
// @Description: 表默认设置
// @param comment 表注释
// @param engine 表引擎
// @return *gorm.DB
func dbSetTableOptions(engine string, comment string, charset string) *gorm.DB {
	//设置表引擎和表注释
	setValue := fmt.Sprintf("ENGINE=%s COMMENT='%s' CHARSET='%s'", engine, comment, charset)

	//fmt.Println(setValue)

	return Gorm.Mysql.Set("gorm:table_options", setValue)
}
