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

func (w *WireGuard) GetUsage() string {
	return `常用命令:

1. 生成密钥对:
   wg genkey | tee privatekey | wg pubkey > publickey

2. 创建配置文件 (/etc/wireguard/wg0.conf):
   [Interface]
   PrivateKey = <your-private-key>
   Address = 10.0.0.1/24
   ListenPort = 51820

   [Peer]
   PublicKey = <peer-public-key>
   AllowedIPs = 10.0.0.2/32

3. 启动 VPN:
   wg-quick up wg0

4. 停止 VPN:
   wg-quick down wg0

5. 查看 VPN 状态:
   wg show

6. 开机自启动:
   systemctl enable wg-quick@wg0

更多信息:
   man wg
   man wg-quick
   https://www.wireguard.com/quickstart/`
}
