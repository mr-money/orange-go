package Config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/goinggo/mapstructure"
	"reflect"
)

var configDir = "./Config"

// Configs 全局配置内容
var Configs struct {
	Web Web
}

//
// Include
// @Description: 初始化配置文件
// @param configs
//
func Include(configs ...interface{}) {
	for _, conf := range configs {
		//反射获取conf文件名
		confRef := reflect.ValueOf(conf)
		confFile := getConfStructName(confRef)

		_, confErr := toml.DecodeFile(confFile, &conf)
		if confErr != nil {
			panic(confErr)
		}

		putConfStruct(confRef, conf)

	}
}

//
//  getConfStructName
//  @Description: 反射获取conf结构体名称
//  @param confRef
//  @return string
//
func getConfStructName(confRef reflect.Value) string {
	return fmt.Sprintf("%v/%v.toml", configDir, confRef.FieldByName("FileName"))
}

//
//  getConfFileName
//  @Description: 获取conf文件名
//  @param confRef
//  @return string
//
func getConfFileName(confRef reflect.Value) string {
	return fmt.Sprintf("%v", confRef.Type())
}

//
//  putConfStruct
//  @Description: 配置map赋值struct
//  @param confRef
//  @param conf
//
func putConfStruct(confRef reflect.Value, conf interface{}) {
	fmt.Println("---------", conf)

	switch getConfFileName(confRef) {
	case "Config.Web": //默认web配置
		_ = mapstructure.Decode(conf, &Configs.Web)

		fmt.Println("---------", Configs.Web)

		break
	}
}
