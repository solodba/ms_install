package impl_test

import "testing"

func TestUploadMysqlInstallFile(t *testing.T) {
	err := svc.UploadMysqlInstallFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
