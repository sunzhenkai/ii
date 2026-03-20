package utils

import (
	"os/exec"
	"runtime"
	"strings"
)

// OSInfo 操作系统信息
type OSInfo struct {
	Type    string // linux, darwin, windows
	Arch    string // amd64, arm64
	Distro  string // ubuntu, centos, etc. (仅 Linux)
	Version string // 系统版本
}

// GetOSInfo 获取操作系统信息
func GetOSInfo() *OSInfo {
	info := &OSInfo{
		Type: runtime.GOOS,
		Arch: runtime.GOARCH,
	}

	// 如果是 Linux，尝试获取发行版信息
	if info.Type == "linux" {
		info.Distro = getLinuxDistro()
	}

	return info
}

// getLinuxDistro 获取 Linux 发行版信息
func getLinuxDistro() string {
	// 尝试读取 /etc/os-release
	if data, err := exec.Command("sh", "-c", "grep '^ID=' /etc/os-release | cut -d= -f2").Output(); err == nil {
		return strings.TrimSpace(strings.ToLower(string(data)))
	}

	// 尝试其他方法
	if _, err := exec.LookPath("apt-get"); err == nil {
		return "debian"
	}
	if _, err := exec.LookPath("yum"); err == nil {
		return "centos"
	}
	if _, err := exec.LookPath("dnf"); err == nil {
		return "fedora"
	}

	return "unknown"
}

// CommandExists 检查命令是否存在
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// RunCommand 运行命令并返回输出
func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// RunCommandInDir 在指定目录运行命令
func RunCommandInDir(dir, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// GetPackageManager 获取系统包管理器
func GetPackageManager() string {
	osInfo := GetOSInfo()

	switch osInfo.Type {
	case "linux":
		switch osInfo.Distro {
		case "ubuntu", "debian", "linuxmint":
			return "apt"
		case "centos", "rhel", "rocky", "almalinux":
			return "yum"
		case "fedora":
			return "dnf"
		case "arch", "manjaro":
			return "pacman"
		case "opensuse", "sles":
			return "zypper"
		}
	case "darwin":
		if CommandExists("brew") {
			return "brew"
		}
	case "windows":
		if CommandExists("winget") {
			return "winget"
		}
		if CommandExists("choco") {
			return "choco"
		}
	}

	return ""
}

// PackageManagerExists 检查指定的包管理器是否存在
func PackageManagerExists(pm string) bool {
	switch pm {
	case "apt":
		return CommandExists("apt-get")
	case "yum":
		return CommandExists("yum")
	case "dnf":
		return CommandExists("dnf")
	case "pacman":
		return CommandExists("pacman")
	case "zypper":
		return CommandExists("zypper")
	case "brew":
		return CommandExists("brew")
	case "winget":
		return CommandExists("winget")
	case "choco":
		return CommandExists("choco")
	default:
		return false
	}
}
