package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/master"
	"github.com/solodba/ms_install/test/tools"
)

var (
	svc master.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(master.AppName).(master.Service)
}
