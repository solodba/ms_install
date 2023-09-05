package protocol

import (
	"fmt"

	"github.com/solodba/ms_install/conf"
)

// Slavea安装程序
func (m *MsInstallSvc) SlaveaInstall() error {
	err := m.slaveaSvc.StopFirewall(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.StopSelinux(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.UploadMysqlInstallFile(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.UnzipMySQLFile(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.CreateMySQLDir(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.IsMySQLRun(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.CreateMySQLUser(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.ChangeMySQLDirPerm(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.InitialMySQL(ctx)
	if err != nil {
		return err
	}
	err = m.slaveaSvc.StartMySQL(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Slavea主从配置
func (m *MsInstallSvc) MsSlaveaInstall() error {
	installType := conf.C().MySQL.InstallType
	switch installType {
	case "pos":
		err := m.slaveaSvc.CloseGtid(ctx)
		if err != nil {
			return err
		}
		err = m.slaveaSvc.ImportFullData(ctx)
		if err != nil {
			return err
		}
		err = m.slaveaSvc.SyncMasterData(ctx)
		if err != nil {
			return err
		}
	case "gtid":
		err := m.slaveaSvc.ImportFullData(ctx)
		if err != nil {
			return err
		}
		err = m.slaveaSvc.SyncMasterGtidData(ctx)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("该安装类型不支持! 目前支持类型: pos(基于位点复制), gtid(基于gtid复制)")
	}
	return nil
}
