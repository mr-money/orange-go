package Config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"reflect"
)

var configDir = "./Config"

// Configs 全局配置内容
var Configs = []interface{}{}

//
// Include
// @Description: 初始化配置文件
// @param configs
//
func Include(configs ...interface{}) {
	for _, conf := range configs {
		//反射获取conf文件名
		ds := reflect.ValueOf(conf)
		confFile := fmt.Sprintf("%v/%v.toml", configDir, ds.FieldByName("FileName"))

		_, confErr := toml.DecodeFile(confFile, &conf)

		if confErr != nil {
			panic(confErr.Error())
			return
		}

		Configs = append(Configs, conf)
	}
}
