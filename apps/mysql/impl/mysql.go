package impl

import (
	"context"
	"fmt"

	"github.com/solodba/mcube/logger"
)

// 上传mysql安装文件
func (i *impl) UploadMysqlInstallFile(context.Context) error {
	srcFile := "mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz"
	dstFile := fmt.Sprintf("/tmp/%s", i.c.MySQL.FileName)
	// master节点
	uploadMsg, err := i.c.Master.UploadFile(srcFile, dstFile)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("[%s]%s", i.c.Master.SysHost, uploadMsg)

	// slavea节点
	uploadMsg, err = i.c.Slavea.UploadFile(srcFile, dstFile)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("[%s]%s", i.c.Slavea.SysHost, uploadMsg)

	// slavea节点
	uploadMsg, err = i.c.Slaveb.UploadFile(srcFile, dstFile)
	if err != nil {
		return err
	}
	logger.L().Info().Msgf("[%s]%s", i.c.Slaveb.SysHost, uploadMsg)
	return nil
}

// 解压MySQL压缩文件
func (i *impl) UnzipMySQLFile(context.Context) error {
	return nil
}

// 创建MySQL相关目录
func (i *impl) CreateMySQLDir(context.Context) error {
	return nil
}

// 判断是否有MySQL进程
func (i *impl) IsMySQLRun(context.Context) error {
	return nil
}

// 创建MySQL用户
func (i *impl) CreateMySQLUser(context.Context) error {
	return nil
}

// 修改权限
func (i *impl) ChangeMySQLDirPerm(context.Context) error {
	return nil
}

// MySQL初始化
func (i *impl) InitialMySQL(context.Context) error {
	return nil
}

// 启动MySQL
func (i *impl) StartMySQL(context.Context) error {
	return nil
}

// 增加环境量变量
func (i *impl) AddEnv(context.Context) error {
	return nil
}
