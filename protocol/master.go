package protocol

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
