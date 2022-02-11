package Config

type Web struct {
	FileName string
	Common   interface{} `toml:"COMMON"`
	DB       map[string]struct {
		Host    string `toml:"host"`
		Port    uint16 `toml:"port"`
		DbName  string `toml:"db_name"`
		User    string `toml:"user"`
		Pwd     string `toml:"pwd"`
		Prefix  string `toml:"prefix"`
		Charest string `toml:"charest"`
		Type    string `toml:"type"`
	} `toml:"DB"`
	Redis map[string]struct {
		Host   string `toml:"host"`
		Port   uint16 `toml:"port"`
		Db     uint8  `toml:"db"`
		Pwd    string `toml:"pwd"`
		Prefix string `toml:"prefix"`
	} `toml:"REDIS"`
}
