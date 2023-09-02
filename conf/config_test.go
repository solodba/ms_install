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
