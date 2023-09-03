package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

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
func (i *impl) StopSelinux(ctx context.Context) error {
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
func (i *impl) UploadMysqlInstallFile(ctx context.Context) error {
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
func (i *impl) UnzipMySQLFile(ctx context.Context) error {
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
	tarName := strings.TrimRight(i.c.MySQL.FileName, ".xz")
	cmd = fmt.Sprintf(`/bin/sh -c "tar xf /tmp/%s -C /tmp"`, tarName)
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("[%s]主机上解压文件[%s]失败, 原因: %s", i.c.Master.SysHost, tarName, err.Error())
	}
	logger.L().Info().Msgf("[%s]主机上解压文件[%s]成功", i.c.Master.SysHost, tarName)
	dirName := strings.TrimRight(tarName, ".tar")
	cmd = fmt.Sprintf(`/bin/sh -c "mv /tmp/%s %s/"`, dirName, i.c.MySQL.InstallPath)
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("[%s]主机上mv文件[%s]失败, 原因: %s", i.c.Master.SysHost, dirName, err.Error())
	}
	logger.L().Info().Msgf("[%s]主机上mv文件[%s]成功", i.c.Master.SysHost, dirName)
	return nil
}

// 创建目录
func (i *impl) CreateDir(ctx context.Context, hostName, dirName string) error {
	cmd := fmt.Sprintf(`/bin/sh -c "mkdir -p %s;chmod 755 -R %s"`, dirName, dirName)
	_, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("[%s]主机上目录[%s]创建失败, 原因: %s", hostName, dirName, err.Error())
	}
	logger.L().Info().Msgf("[%s]主机上目录[%s]创建成功", hostName, dirName)
	return nil
}

// 创建MySQL相关目录
func (i *impl) CreateMySQLDir(ctx context.Context) error {
	cmd := fmt.Sprintf(`/bin/sh -c "ls -ld %s"`, i.c.MySQL.BaseDir)
	_, err := i.c.Master.RunShell(cmd)
	time.Sleep(3 * time.Second)
	if err == nil {
		return fmt.Errorf("[%s]文件夹[%s]已经存在, 请确定是否安装了MySQL", i.c.Master.SysHost, i.c.MySQL.BaseDir)
	}
	// 创建数据目录
	err = i.CreateDir(ctx, i.c.Master.SysHost, i.c.MySQL.DataPath())
	if err != nil {
		return err
	}
	// 创建binlog目录
	err = i.CreateDir(ctx, i.c.Master.SysHost, i.c.MySQL.BinlogPath())
	if err != nil {
		return err
	}
	// 创建日志目录
	err = i.CreateDir(ctx, i.c.Master.SysHost, i.c.MySQL.LogPath())
	if err != nil {
		return err
	}
	// 创建临时文件目录
	err = i.CreateDir(ctx, i.c.Master.SysHost, i.c.MySQL.TmpPath())
	if err != nil {
		return err
	}
	// 创建配置文件目录
	err = i.CreateDir(ctx, i.c.Master.SysHost, i.c.MySQL.ConfPath())
	if err != nil {
		return err
	}
	// 创建备份文件目录
	err = i.CreateDir(ctx, i.c.Master.SysHost, i.c.MySQL.BackupPath())
	if err != nil {
		return err
	}
	return nil
}

// 判断是否有MySQL进程
func (i *impl) IsMySQLRun(ctx context.Context) error {
	cmd := fmt.Sprintf(`/bin/sh -c "ps -ef | grep mysqld | grep -v grep | wc -l"`)
	res, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
	}
	if strings.Trim(res, "\n") != "0" {
		return fmt.Errorf("有MySQL进程在运行, 请检查")
	}
	return nil
}

// 创建MySQL用户
func (i *impl) CreateMySQLUser(context.Context) error {
	cmd := fmt.Sprintf(`/bin/sh -c "cat /etc/passwd |grep -w mysql|wc -l"`)
	res, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
	}
	if strings.Trim(res, "\n") != "1" {
		cmd = fmt.Sprintf(`/bin/sh -c "groupadd mysql"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
		}
		cmd = fmt.Sprintf(`/bin/sh -c "useradd -g mysql mysql"`)
		_, err = i.c.Master.RunShell(cmd)
		if err != nil {
			return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
		}
		logger.L().Info().Msgf("添加mysql用户成功")
	}
	return nil
}

// 修改权限
func (i *impl) ChangeMySQLDirPerm(context.Context) error {
	cmd := fmt.Sprintf(`/bin/sh -c "chown -R mysql:mysql %s"`, i.c.MySQL.BaseDir)
	_, err := i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
	}
	cmd = fmt.Sprintf(`/bin/sh -c "chown -R mysql:mysql %s"`, i.c.MySQL.InstallPath)
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
	}
	cmd = fmt.Sprintf(`/bin/sh -c "cp my.cnf %s/"`, i.c.MySQL.ConfPath())
	_, err = i.c.Master.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("执行命令[%s]报错, 原因: %s", cmd, err.Error())
	}
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
