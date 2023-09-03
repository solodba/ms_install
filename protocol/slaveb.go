package protocol

// Master安装程序
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
