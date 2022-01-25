package Config

import (
	"github.com/BurntSushi/toml"
)

var configDir = "./Config/"

var configs = []interface{}{}

func Include(conf ...interface{}) {
	configs = append(configs, conf...)
}

func Init() {
	var webConfig Web

	//读取配置
	_, err := toml.DecodeFile(configDir+"web.toml", &webConfig)
	if err != nil {
		//fmt.Println(err)
		panic(err)
		return
	}

}
