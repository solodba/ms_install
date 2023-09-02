package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/ms_install/apps/slaveb"
	"github.com/solodba/ms_install/test/tools"
)

var (
	svc slaveb.Service
	ctx = context.Background()
)

func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(slaveb.AppName).(slaveb.Service)
}
