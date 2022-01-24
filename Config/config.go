package Config

type WebConfig struct {
	Common interface{}
	DB     struct {
		Local struct {
			Host    string
			Port    uint16
			DbName  string
			User    string
			Pwd     string
			Prefix  string
			Charest string
			Type    string
		}
	}
	Redis struct {
		Local struct {
			Host   string
			Port   uint16
			Db     uint8
			Pwd    string
			Prefix string
		}
	}
}

func ConfInit() {
	//https://github.com/BurntSushi/toml
}
