package protocol

import (
	"context"
	"fmt"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/master"
	"github.com/solodba/ms_install/apps/slavea"
	"github.com/solodba/ms_install/apps/slaveb"
	"github.com/solodba/ms_install/conf"
)

var (
	ctx = context.Background()
)

// 主从安装服务结构体
type MsInstallSvc struct {
	masterSvc master.Service
	slaveaSvc slavea.Service
	slavebSvc slaveb.Service
}

// MySQL主从安装服务结构体构造函数
func NewMsInstallSvc() *MsInstallSvc {
	return &MsInstallSvc{
		masterSvc: apps.GetInternalApp(master.AppName).(master.Service),
		slaveaSvc: apps.GetInternalApp(slavea.AppName).(slavea.Service),
		slavebSvc: apps.GetInternalApp(slaveb.AppName).(slaveb.Service),
	}
}

// MySQL安装服务
func (m *MsInstallSvc) MsInstall(ctx context.Context) error {
	mode := conf.C().MySQL.InstallMode
	switch mode {
	case "mo":
		// 只安装master节点
		err := m.MasterInstall()
		if err != nil {
			return err
		}
	case "sao":
		// 只安装slavea节点
		err := m.SlaveaInstall()
		if err != nil {
			return err
		}
	case "sbo":
		// 只安装slaveb节点
		err := m.SlavebInstall()
		if err != nil {
			return err
		}
	case "ms":
		// 安装一主一从
		err := m.MasterInstall()
		if err != nil {
			return err
		}
		// 只安装slavea节点
		err = m.SlaveaInstall()
		if err != nil {
			return err
		}
		// 主节点主从配置
		err = m.MsMasterInstall()
		if err != nil {
			return err
		}
		// slavea节点主从配置
		err = m.MsSlaveaInstall()
		if err != nil {
			return err
		}
	case "mss":
		// 安装一主两从
		err := m.MasterInstall()
		if err != nil {
			return err
		}
		// 安装slavea节点
		err = m.SlaveaInstall()
		if err != nil {
			return err
		}
		// 安装slaveb节点
		err = m.SlavebInstall()
		if err != nil {
			return err
		}
		// 主节点主从配置
		err = m.MsMasterInstall()
		if err != nil {
			return err
		}
		// slavea节点主从配置
		err = m.MsSlaveaInstall()
		if err != nil {
			return err
		}
		// slaveb节点主从配置
		err = m.MsSlavebInstall()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("该安装类型不支持! 目前支持类型: mo(只安装master节点), sao(只安装slavea节点), sbo(只安装slaveb节点), ms(安装一主一从), mss(安装一主两从)")
	}
	return nil
}
