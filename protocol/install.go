package protocol

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/master"
	"github.com/solodba/ms_install/apps/slavea"
	"github.com/solodba/ms_install/apps/slaveb"
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
	// 主节点安装MySQL
	// err := m.MasterInstall()
	// if err != nil {
	// 	return err
	// }

	// 从节点A安装MySQL
	// err := m.SlaveaInstall()
	// if err != nil {
	// 	return err
	// }

	// 从节点B安装MySQL
	err := m.SlavebInstall()
	if err != nil {
		return err
	}
	return nil
}
