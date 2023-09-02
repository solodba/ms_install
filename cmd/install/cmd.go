package install

import (
	"context"

	"github.com/solodba/ms_install/protocol"
	"github.com/spf13/cobra"
)

var (
	ctx = context.Background()
)

// MySQL服务结构体
type Server struct {
	MySQLInstallSvc *protocol.MsInstallSvc
}

// MySQL服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		MySQLInstallSvc: protocol.NewMsInstallSvc(),
	}
}

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "install",
	Short:   "ms_install install service",
	Long:    "ms_install install service",
	Example: "ms_install install -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := NewServer()
		err := svc.MySQLInstallSvc.MsInstall(ctx)
		if err != nil {
			return err
		}
		return nil
	},
}
