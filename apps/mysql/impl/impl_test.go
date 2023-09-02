package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/mysql"
	"github.com/solodba/ms_install/test/tools"
)

var (
	svc mysql.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(mysql.AppName).(mysql.Service)
}
