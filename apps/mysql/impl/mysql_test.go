package impl_test

import "testing"

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
