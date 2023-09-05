package conf

import "fmt"

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
	BackupDirName string `toml:"backup_dir_name" env:"MYSQL_BACKUP_DIR_NAME"`
	RootPassword  string `toml:"root_password" env:"MYSQL_ROOT_PASSWORD"`
	ReplUser      string `toml:"repl_user" env:"MYSQL_REPL_USER"`
	ReplPassword  string `toml:"repl_password" env:"MYSQL_REPL_PASSWORD"`
	InstallMode   string `toml:"install_mode" env:"MYSQL_INSTALL_MODE"`
	InstallType   string `toml:"install_type" env:"MYSQL_INSTALL_TYPE"`
}

// Master结构体
type Master struct {
	SysUsername string `toml:"sys_username" env:"MASTER_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"MASTER_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"MASTER_SYS_HOST"`
	SysPort     int64  `toml:"sys_port" env:"MASTER_SYS_PORT"`
	ServerId    int64  `toml:"server_id" env:"MASTER_SERVER_ID"`
}

// Slavea结构体
type Slavea struct {
	SysUsername string `toml:"sys_username" env:"SLAVEA_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"SLAVEA_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"SLAVEA_SYS_HOST"`
	SysPort     int64  `toml:"sys_port" env:"SLAVEA_SYS_PORT"`
	ServerId    int64  `toml:"server_id" env:"SLAVEA_SERVER_ID"`
}

// Slaveb结构体
type Slaveb struct {
	SysUsername string `toml:"sys_username" env:"SLAVEB_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"SLAVEB_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"SLAVEB_SYS_HOST"`
	SysPort     int64  `toml:"sys_port" env:"SLAVEB_SYS_PORT"`
	ServerId    int64  `toml:"server_id" env:"SLAVEB_SERVER_ID"`
}

// Config构造函数
func NewDefaultConfig() *Config {
	return &Config{
		MySQL:  NewDefaultMySQL(),
		Master: NewDefaultMaster(),
		Slavea: NewDefaultSlavea(),
		Slaveb: NewDefaultSlaveb(),
	}
}

// MySQL结构体构造函数
func NewDefaultMySQL() *MySQL {
	return &MySQL{}
}

// Master构造函数
func NewDefaultMaster() *Master {
	return &Master{}
}

// Slavea构造函数
func NewDefaultSlavea() *Slavea {
	return &Slavea{}
}

// Slaveb构造函数
func NewDefaultSlaveb() *Slaveb {
	return &Slaveb{}
}

// 获取MySQL安装路径
func (m *MySQL) BinlogPath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.BinlogDirName)
}

func (m *MySQL) DataPath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.DataDirName)
}

func (m *MySQL) LogPath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.LogDirName)
}

func (m *MySQL) TmpPath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.TmpDirName)
}

func (m *MySQL) ConfPath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.ConfDirName)
}

func (m *MySQL) BackupPath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.BackupDirName)
}
