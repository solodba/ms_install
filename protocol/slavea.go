package protocol

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
	return nil
}
