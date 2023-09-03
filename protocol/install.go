package protocol

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/master"
)

var (
	ctx = context.Background()
)

// 主从安装服务结构体
type MsInstallSvc struct {
	masterSvc master.Service
}

// MySQL主从安装服务结构体构造函数
func NewMsInstallSvc() *MsInstallSvc {
	return &MsInstallSvc{
		masterSvc: apps.GetInternalApp(master.AppName).(master.Service),
	}
}

// Master安装程序
func (m *MsInstallSvc) MasterInstall() error {
	err := m.masterSvc.StopFirewall(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.StopSelinux(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.UploadMysqlInstallFile(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.UnzipMySQLFile(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.CreateMySQLDir(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.IsMySQLRun(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.CreateMySQLUser(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.ChangeMySQLDirPerm(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.InitialMySQL(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.StartMySQL(ctx)
	if err != nil {
		return err
	}
	return nil
}

// MySQL安装服务
func (m *MsInstallSvc) MsInstall(ctx context.Context) error {
	err := m.MasterInstall()
	if err != nil {
		return err
	}
	return nil
}
