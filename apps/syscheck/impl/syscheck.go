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
	_, err = i.c.Slavea.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("主机[%s]防火墙关闭失败, 原因: %s", i.c.Slavea.SysHost, err.Error())
	}
	logger.L().Info().Msgf("主机[%s]防火墙关闭成功", i.c.Slavea.SysHost)
	_, err = i.c.Slaveb.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("主机[%s]防火墙关闭失败, 原因: %s", i.c.Slaveb.SysHost, err.Error())
	}
	logger.L().Info().Msgf("主机[%s]防火墙关闭成功", i.c.Slaveb.SysHost)
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

	slaveaFlag, err := i.c.Slavea.RunShell(`/bin/bash -c "getenforce"`)
	if strings.TrimRight(slaveaFlag, "\n") != "Disabled" {
		_, err = i.c.Slavea.RunShell(cmd)
		if err != nil {
			return fmt.Errorf("主机[%s]selinux关闭失败, 原因: %s", i.c.Slavea.SysHost, err.Error())
		}
		logger.L().Info().Msgf("主机[%s]Selinux关闭成功", i.c.Slavea.SysHost)
	} else {
		logger.L().Info().Msgf("主机[%s]Selinux已经关闭", i.c.Slavea.SysHost)
	}

	slavebFlag, err := i.c.Slaveb.RunShell(`/bin/bash -c "getenforce"`)
	if strings.TrimRight(slavebFlag, "\n") != "Disabled" {
		_, err = i.c.Slaveb.RunShell(cmd)
		if err != nil {
			return fmt.Errorf("主机[%s]selinux关闭失败, 原因: %s", i.c.Slaveb.SysHost, err.Error())
		}
		logger.L().Info().Msgf("主机[%s]Selinux关闭成功", i.c.Slaveb.SysHost)
	} else {
		logger.L().Info().Msgf("主机[%s]Selinux已经关闭", i.c.Slaveb.SysHost)
	}

	return nil
}
