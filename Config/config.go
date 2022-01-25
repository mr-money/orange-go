package Config

import (
	"github.com/BurntSushi/toml"
)

//var configFile = map[string]string{
//	"Web": "web.toml",
//}

var _ConfigDir_ = "./Config/"

func Init() {
	var webConfig Web

	//读取配置
	_, err := toml.DecodeFile(_ConfigDir_+"web.toml", &webConfig)
	if err != nil {
		//fmt.Println(err)
		panic(err)
		return
	}

}
