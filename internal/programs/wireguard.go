package programs

import (
	"github.com/wii/ii/pkg/types"
)

// WireGuard 程序定义
type WireGuard struct{}

func NewWireGuard() types.Program {
	return &WireGuard{}
}

func (w *WireGuard) Name() string {
	return "wireguard"
}

func (w *WireGuard) Description() string {
	return "WireGuard 是一种极其简单、快速、现代的 VPN 技术"
}

func (w *WireGuard) GetInstallMethods() map[string]string {
	// 返回支持的安装方法和对应的包名
	// 格式: map[安装方法名]包名
	return map[string]string{
		// Linux 系统包管理器
		"apt":    "wireguard",
		"yum":    "wireguard-tools",
		"dnf":    "wireguard-tools",
		"pacman": "wireguard-tools",
		"zypper": "wireguard-tools",

		// macOS/Linux
		"brew": "wireguard-tools",

		// 版本管理器（如果 WireGuard 有相关插件）
		// 注：WireGuard 通常通过系统包管理器安装，这里只是示例
		// "mise": "wireguard",
		// "asdf": "wireguard",
	}
}

func (w *WireGuard) GetSupportedPlatforms() map[string][]string {
	return map[string][]string{
		"linux":  {"amd64", "arm64", "arm"},
		"darwin": {"amd64", "arm64"},
	}
}
