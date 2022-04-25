package Config

type Web struct {
	FileName string
	Common   interface{} `toml:"COMMON"`
	DB       struct {
		Host    string `toml:"host"`
		Port    string `toml:"port"`
		DbName  string `toml:"dbName"`
		User    string `toml:"user"`
		Pwd     string `toml:"pwd"`
		Prefix  string `toml:"prefix"`
		Charset string `toml:"charset"`
		Type    string `toml:"type"`
	} `toml:"DB"`
	Redis struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Db   string `toml:"db"`
		Pwd  string `toml:"pwd"`
	} `toml:"REDIS"`
}
