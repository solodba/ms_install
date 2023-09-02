package syscheck

import "context"

// 模块名称
const (
	AppName = "syscheck"
)

// 关闭系统服务接口
type Service interface {
	// 关闭防火墙
	StopFirewall(context.Context) error
	// 关闭selinux
	StopSelinux(context.Context) error
}
