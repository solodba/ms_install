package slavea

import "context"

// 模块名称
const (
	AppName = "slavea"
)

// 从节点A安装服务接口
type Service interface {
	// 关闭防火墙
	StopFirewall(context.Context) error
	// 关闭selinux
	StopSelinux(context.Context) error
	// 上传mysql安装文件
	UploadMysqlInstallFile(context.Context) error
	// 解压MySQL压缩文件
	UnzipMySQLFile(context.Context) error
	// 创建MySQL相关目录
	CreateMySQLDir(context.Context) error
	// 判断是否有MySQL进程
	IsMySQLRun(context.Context) error
	// 创建MySQL用户
	CreateMySQLUser(context.Context) error
	// 修改权限
	ChangeMySQLDirPerm(context.Context) error
	// MySQL初始化
	InitialMySQL(context.Context) error
	// 启动MySQL
	StartMySQL(context.Context) error
	// 增加环境量变量
	AddEnv(context.Context) error
	// 关闭GTID
	CloseGtid(context.Context) error
	// 全库数据导入
	ImportFullData(context.Context) error
	// 获取binlogfile和position
	GetBinLogFileNameAndPos(context.Context) (*BinLogFileNamePos, error)
	// 从库配置同步
	SyncMasterData(context.Context) error
	// 从库配置GTID数据同步
	SyncMasterGtidData(context.Context) error
}
