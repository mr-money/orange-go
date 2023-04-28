package Config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/goinggo/mapstructure"
	"github.com/shockerli/cvt"
	"go-study/Library/Handler"
	"log"
	"reflect"
)

// Configs 全局配置内容
var Configs struct {
	Web Web
}

func init() {
	//加载配置
	var webConfig Web
	webConfig.FileName = "web"

	include(webConfig)
}

//
// Include
// @Description: 初始化配置文件
// @param configs
//
func include(configs ...interface{}) {
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
	rootPath, _ := Handler.GetProjectRoot()

	return fmt.Sprintf("%v/Config/%v.toml", rootPath, confRef.FieldByName("FileName"))
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
	//fmt.Println("---------", conf)

	switch getConfFileName(confRef) {
	case "Config.Web": //默认web配置
		//fmt.Println("---------", conf)

		err := mapstructure.Decode(conf, &Configs.Web)
		if err != nil {
			log.Panicln(err)
		}

		//fmt.Println("---------", Configs.Web)

		break
	}
}

//// 公共方法 ////

// GetFieldByName
// @Description: 反射获取配置值
// @param confStruct 配置结构体
// @param fieldName 结构体内字段名 如 DB,Host
// @return string
//
func GetFieldByName(confStruct interface{}, fieldName ...string) string {
	var fieldNames []string
	fieldNames = append(fieldNames, fieldName...)

	conf := reflect.ValueOf(confStruct)
	for _, field := range fieldNames {
		conf = conf.FieldByName(field)
	}

	return cvt.String(conf)
}
