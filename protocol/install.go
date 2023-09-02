package protocol

import (
	"context"
)

// 主从安装服务结构体
type MsInstallSvc struct {
}

// MySQL主从安装服务结构体构造函数
func NewMsInstallSvc() *MsInstallSvc {
	return nil
}

// MySQL安装服务
func (m *MsInstallSvc) MsInstall(ctx context.Context) error {
	return nil
}
