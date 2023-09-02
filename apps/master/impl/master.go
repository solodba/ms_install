package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/solodba/mcube/logger"
)

// 关闭防火墙
func (i *impl) StopFirewall(ctx context.Context) error {
	cmd := `/bin/bash -c "systemctl stop firewalld;systemctl disable firewalld;systemctl mask firewalld"`
	_, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("主机[%s]防火墙关闭失败, 原因: %s", i.c.Master.SysHost, err.Error())
	}
	logger.L().Info().Msgf("主机[%s]防火墙关闭成功", i.c.Master.SysHost)
	return nil
}

// 关闭selinux
func (i *impl) StopSelinux(context.Context) error {
	cmd := `/bin/bash -c "sed -i 's/enforcing/disabled/' /etc/selinux/config;setenforce 0"`
	masterFlag, err := i.c.Master.RunShell(`/bin/bash -c "getenforce"`)
	if err != nil {
		return err
	}
	if strings.TrimRight(masterFlag, "\n") != "Disabled" {
		_, err := i.c.Master.RunShell(cmd)
		if err != nil {
			return fmt.Errorf("主机[%s]selinux关闭失败, 原因: %s", i.c.Master.SysHost, err.Error())
		}
		logger.L().Info().Msgf("主机[%s]Selinux关闭成功", i.c.Master.SysHost)
	} else {
		logger.L().Info().Msgf("主机[%s]Selinux已经关闭", i.c.Master.SysHost)
	}
	return nil
}

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
	return nil
}

// 解压MySQL压缩文件
func (i *impl) UnzipMySQLFile(context.Context) error {
	remoteFile := fmt.Sprintf("/tmp/%s", i.c.MySQL.FileName)
	// master节点
	cmd := fmt.Sprintf(`ls -ld %s`, i.c.MySQL.InstallPath)
	_, err := i.c.Master.RunShell(cmd)
	if err == nil {
		return fmt.Errorf("[%s]主机上[%s]文件夹已经存在,请确定是否安装了MySQL", i.c.Master.SysHost, i.c.MySQL.InstallPath)
	}
	cmd = fmt.Sprintf(`/bin/sh -c "xz -d %s"`, remoteFile)
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("[%s]主机上解压文件[%s]失败, 原因: %s", i.c.Master.SysHost, remoteFile, err.Error())
	}
	logger.L().Info().Msgf("[%s]主机上解压文件[%s]成功", i.c.Master.SysHost, remoteFile)
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
