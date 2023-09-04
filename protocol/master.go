package protocol

import (
	"fmt"

	"github.com/solodba/ms_install/conf"
)

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

// Master上主从配置
func (m *MsInstallSvc) MsMasterInstall() error {
	mode := conf.C().MySQL.InstallMode
	err := m.masterSvc.CloseGtid(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.CreateReplicateUser(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.MySqlPosDataDump(ctx)
	if err != nil {
		return err
	}
	err = m.masterSvc.DownLoadPosDataFile(ctx)
	if err != nil {
		return err
	}
	switch mode {
	case "ms":
		err = m.masterSvc.CopyDumpDataToSlavea(ctx)
		if err != nil {
			return err
		}
	case "mss":
		err = m.masterSvc.CopyDumpDataToSlavea(ctx)
		if err != nil {
			return err
		}
		err = m.masterSvc.CopyDumpDataToSlaveb(ctx)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("该安装类型不支持! 目前支持类型: mo(只安装master节点), sao(只安装slavea节点), sbo(只安装slaveb节点), ms(安装一主一从), mss(安装一主两从)")
	}
	return nil
}
