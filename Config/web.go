package Config

type Web struct {
	Common interface{}      `toml:"COMMON"`
	DB     map[string]db    `toml:"DB"`
	Redis  map[string]redis `toml:"REDIS"`
}

type db struct {
	Host    string `toml:"host"`
	Port    uint16 `toml:"port"`
	DbName  string `toml:"db_name"`
	User    string `toml:"user"`
	Pwd     string `toml:"pwd"`
	Prefix  string `toml:"prefix"`
	Charest string `toml:"charest"`
	Type    string `toml:"type"`
}

type redis struct {
	Host   string `toml:"host"`
	Port   uint16 `toml:"port"`
	Db     uint8  `toml:"db"`
	Pwd    string `toml:"pwd"`
	Prefix string `toml:"prefix"`
}
