package Database

import "go-study/Model"

var mysqlMigrations []map[string]interface{}

//
// getMysqlMigrations
// @Description: mysql连接 迁移ORM
// @return []map[string]interface{}
//
func getMysqlMigrations() []map[string]interface{} {
	return append(mysqlMigrations,
		//mysql下 user 用户表
		map[string]interface{}{
			"engine":  Model.User{}.GetOption("engine"),
			"comment": Model.User{}.GetOption("comment"),
			"charset": Model.User{}.GetOption("charset"),
			"model":   Model.User{},
		},

		//其他表...
		/*map[string]interface{}{
			"engine":  Model.modelName{}.GetOption("engine"),
			"comment": Model.modelName{}.GetOption("comment"),
			"charset": Model.modelName{}.GetOption("charset"),
			"model":   Model.modelName{},
		},*/
	)
}
