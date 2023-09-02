package conf

// 全局配置
var (
	c *Config
)

// 通过函数返回初始化的结构体
func C() *Config {
	if c == nil {
		panic("请初始化全局配置!")
	}
	return c
}

// Config结构体
type Config struct {
	MySQL  *MySQL  `toml:"mysql"`
	Master *Master `toml:"master"`
	Slavea *Slavea `toml:"slavea"`
	Slaveb *Slaveb `toml:"slaveb"`
}

// MySQL结构体
type MySQL struct {
	FileName      string `toml:"file_name" env:"MYSQL_FILE_NAME"`
	InstallPath   string `toml:"install_path" env:"MYSQL_INSTALL_PATH"`
	BaseDir       string `toml:"base_dir" env:"MYSQL_BASE_DIR"`
	BinlogDirName string `toml:"binlog_dir_name" env:"MYSQL_BINLOG_DIR_NAME"`
	DataDirName   string `toml:"data_dir_name" env:"MYSQL_DATA_DIR_NAME"`
	LogDirName    string `toml:"log_dir_name" env:"MYSQL_LOG_DIR_NAME"`
	TmpDirName    string `toml:"tmp_dir_name" env:"MYSQL_TMP_DIR_NAME"`
	ConfDirName   string `toml:"conf_dir_name" env:"MYSQL_CONF_DIR_NAME"`
}

// Master结构体
type Master struct {
	SysUsername string `toml:"sys_username" env:"MASTER_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"MASTER_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"MASTER_SYS_HOST"`
	SysPort     string `toml:"sys_port" env:"MASTER_SYS_PORT"`
}

// Slavea结构体
type Slavea struct {
	SysUsername string `toml:"sys_username" env:"SLAVEA_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"SLAVEA_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"SLAVEA_SYS_HOST"`
	SysPort     string `toml:"sys_port" env:"SLAVEA_SYS_PORT"`
}

// Slaveb结构体
type Slaveb struct {
	SysUsername string `toml:"sys_username" env:"SLAVEB_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"SLAVEB_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"SLAVEB_SYS_HOST"`
	SysPort     string `toml:"sys_port" env:"SLAVEB_SYS_PORT"`
}

// Config构造函数
func NewConfig() *Config {
	return &Config{
		MySQL:  NewMySQL(),
		Slavea: NewSlavea(),
		Slaveb: NewSlaveb(),
	}
}

// MySQL结构体构造函数
func NewMySQL() *MySQL {
	return &MySQL{}
}

// Master构造函数
func NewMaster() *Master {
	return &Master{}
}

// Slavea构造函数
func NewSlavea() *Slavea {
	return &Slavea{}
}

// Slaveb构造函数
func NewSlaveb() *Slaveb {
	return &Slaveb{}
}
