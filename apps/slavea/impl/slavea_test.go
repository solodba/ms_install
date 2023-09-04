package impl_test

import "testing"

func TestStopFirewall(t *testing.T) {
	err := svc.StopFirewall(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStopSelinux(t *testing.T) {
	err := svc.StopSelinux(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadMysqlInstallFile(t *testing.T) {
	err := svc.UploadMysqlInstallFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnzipMySQLFile(t *testing.T) {
	err := svc.UnzipMySQLFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCloseGtid(t *testing.T) {
	err := svc.CloseGtid(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImportFullData(t *testing.T) {
	err := svc.ImportFullData(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBinLogFileNameAndPos(t *testing.T) {
	binLogNameAndPos, err := svc.GetBinLogFileNameAndPos(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(binLogNameAndPos)
}

func TestSyncMasterData(t *testing.T) {
	err := svc.SyncMasterData(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
