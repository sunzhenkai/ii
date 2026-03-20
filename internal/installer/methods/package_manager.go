package methods

import (
	"context"
	"fmt"

	"github.com/wii/ii/internal/utils"
	"github.com/wii/ii/pkg/types"
)

// PackageManagerMethod 系统包管理器安装方法
type PackageManagerMethod struct {
	name string
	pm   string
}

// NewPackageManagerMethod 创建包管理器方法
func NewPackageManagerMethod(pm string) types.InstallMethod {
	return &PackageManagerMethod{
		name: pm,
		pm:   pm,
	}
}

func (m *PackageManagerMethod) Name() string {
	return m.name
}

func (m *PackageManagerMethod) Description() string {
	switch m.pm {
	case "apt":
		return "APT (Debian/Ubuntu 包管理器)"
	case "yum":
		return "YUM (CentOS/RHEL 包管理器)"
	case "dnf":
		return "DNF (Fedora 包管理器)"
	case "pacman":
		return "Pacman (Arch Linux 包管理器)"
	case "zypper":
		return "Zypper (openSUSE 包管理器)"
	default:
		return fmt.Sprintf("%s 包管理器", m.pm)
	}
}

func (m *PackageManagerMethod) IsAvailable() bool {
	return utils.PackageManagerExists(m.pm)
}

func (m *PackageManagerMethod) Install(ctx context.Context, program, packageName string) error {
	var cmd string
	var args []string

	switch m.pm {
	case "apt":
		cmd = "sudo"
		args = []string{"apt-get", "install", "-y", packageName}
	case "yum":
		cmd = "sudo"
		args = []string{"yum", "install", "-y", packageName}
	case "dnf":
		cmd = "sudo"
		args = []string{"dnf", "install", "-y", packageName}
	case "pacman":
		cmd = "sudo"
		args = []string{"pacman", "-S", "--noconfirm", packageName}
	case "zypper":
		cmd = "sudo"
		args = []string{"zypper", "install", "-y", packageName}
	default:
		return fmt.Errorf("不支持的包管理器: %s", m.pm)
	}

	output, err := utils.RunCommand(cmd, args...)
	if err != nil {
		return fmt.Errorf("安装失败: %w\n输出: %s", err, output)
	}

	return nil
}

func (m *PackageManagerMethod) GetInstallInfo(program, packageName string) string {
	return fmt.Sprintf("使用 %s 安装 %s (包名: %s)", m.Description(), program, packageName)
}
