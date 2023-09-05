package protocol

import (
	"fmt"

	"github.com/solodba/ms_install/conf"
)

// Slaveb安装程序
func (m *MsInstallSvc) SlavebInstall() error {
	err := m.slavebSvc.StopFirewall(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.StopSelinux(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.UploadMysqlInstallFile(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.UnzipMySQLFile(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.CreateMySQLDir(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.IsMySQLRun(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.CreateMySQLUser(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.ChangeMySQLDirPerm(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.InitialMySQL(ctx)
	if err != nil {
		return err
	}
	err = m.slavebSvc.StartMySQL(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Slavea主从配置
func (m *MsInstallSvc) MsSlavebInstall() error {
	installType := conf.C().MySQL.InstallType
	switch installType {
	case "pos":
		err := m.slavebSvc.CloseGtid(ctx)
		if err != nil {
			return err
		}
		err = m.slavebSvc.ImportFullData(ctx)
		if err != nil {
			return err
		}
		err = m.slavebSvc.SyncMasterData(ctx)
		if err != nil {
			return err
		}
	case "gtid":
		err := m.slavebSvc.ImportFullData(ctx)
		if err != nil {
			return err
		}
		err = m.slavebSvc.SyncMasterGtidData(ctx)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("该安装类型不支持! 目前支持类型: pos(基于位点复制), gtid(基于gtid复制)")

	}
	return nil
}
