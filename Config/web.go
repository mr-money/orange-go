package Config

type Web struct {
	FileName string
	Common   struct {
		EnvModel string `mapstructure:"env_mode"`
	} `mapstructure:"COMMON"`
	DB struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		DbName  string `mapstructure:"dbName"`
		User    string `mapstructure:"user"`
		Pwd     string `mapstructure:"pwd"`
		Prefix  string `mapstructure:"prefix"`
		Charset string `mapstructure:"charset"`
		Type    string `mapstructure:"type"`
	} `mapstructure:"DB"`
	Redis struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Db   string `mapstructure:"db"`
		Pwd  string `mapstructure:"pwd"`
	} `mapstructure:"REDIS"`
}
