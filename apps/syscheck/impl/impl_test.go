package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/syscheck"
	"github.com/solodba/ms_install/test/tools"
)

var (
	svc syscheck.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(syscheck.AppName).(syscheck.Service)
}
