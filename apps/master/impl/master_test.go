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

func TestCreateMySQLDir(t *testing.T) {
	err := svc.CreateMySQLDir(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsMySQLRun(t *testing.T) {
	err := svc.IsMySQLRun(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateMySQLUser(t *testing.T) {
	err := svc.CreateMySQLUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChangeMySQLDirPerm(t *testing.T) {
	err := svc.ChangeMySQLDirPerm(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
