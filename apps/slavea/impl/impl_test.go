package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/slavea"
	"github.com/solodba/ms_install/test/tools"
)

var (
	svc slavea.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(slavea.AppName).(slavea.Service)
}
