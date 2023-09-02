package conf_test

import (
	"testing"

	"github.com/solodba/ms_install/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().MySQL)
	t.Log(conf.C().Master)
	t.Log(conf.C().Slavea)
	t.Log(conf.C().Slaveb)
	t.Log(conf.C().MySQL.DataPath())
	t.Log(conf.C().MySQL.BinlogPath())
	t.Log(conf.C().MySQL.LogPath())
	t.Log(conf.C().MySQL.TmpPath())
	t.Log(conf.C().MySQL.ConfPath())
	t.Log(conf.C().MySQL.BackupPath())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().MySQL)
	t.Log(conf.C().Master)
	t.Log(conf.C().Slavea)
	t.Log(conf.C().Slaveb)
	t.Log(conf.C().MySQL.DataPath())
	t.Log(conf.C().MySQL.BinlogPath())
	t.Log(conf.C().MySQL.LogPath())
	t.Log(conf.C().MySQL.TmpPath())
	t.Log(conf.C().MySQL.ConfPath())
	t.Log(conf.C().MySQL.BackupPath())
}

func TestUploadFile(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	srcFile := "mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz"
	dstFile := "/root/mysql-8.0.25-linux-glibc2.12-x86_64.tar.xz"
	// result, err := conf.C().Master.UploadFile(srcFile, dstFile)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(result)

	// result, err := conf.C().Slavea.UploadFile(srcFile, dstFile)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(result)

	result, err := conf.C().Slaveb.UploadFile(srcFile, dstFile)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRunShell(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	cmd := `/bin/bash -c "ls -l"`
	result, err := conf.C().Master.RunShell(cmd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)

	// result, err := conf.C().Slavea.RunShell(cmd)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(result)

	// result, err := conf.C().Slaveb.RunShell(cmd)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(result)
}
