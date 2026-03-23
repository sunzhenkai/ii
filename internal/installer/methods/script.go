package methods

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/wii/ii/internal/utils"
	"github.com/wii/ii/pkg/types"
)

// ScriptMethod 脚本安装方法（用于通过下载并执行脚本安装）
type ScriptMethod struct {
	name        string
	description string
	scriptURL   string
}

// NewAnacondaScriptMethod 创建 Anaconda 脚本安装方法
func NewAnacondaScriptMethod() types.InstallMethod {
	return &ScriptMethod{
		name:        "anaconda-script",
		description: "官方安装脚本",
		scriptURL:   "https://repo.anaconda.com/archive",
	}
}

func (m *ScriptMethod) Name() string {
	return m.name
}

func (m *ScriptMethod) Description() string {
	return m.description
}

func (m *ScriptMethod) IsAvailable() bool {
	// 需要 wget 或 curl
	return utils.CommandExists("wget") || utils.CommandExists("curl")
}

func (m *ScriptMethod) Install(ctx context.Context, program, packageName string) error {
	// 获取系统架构
	arch := runtime.GOARCH
	if arch == "amd64" {
		arch = "x86_64"
	} else if arch == "arm64" {
		arch = "aarch64"
	}

	// 获取最新版本
	// 这里使用一个固定的最新版本，实际应该从 API 获取
	latestVersion := "2024.10-1"

	// 构建下载 URL
	var filename string
	if runtime.GOOS == "linux" {
		if arch == "aarch64" {
			filename = fmt.Sprintf("Anaconda3-%s-Linux-%s.sh", latestVersion, arch)
		} else {
			filename = fmt.Sprintf("Anaconda3-%s-Linux-%s.sh", latestVersion, arch)
		}
	} else if runtime.GOOS == "darwin" {
		if arch == "arm64" {
			filename = fmt.Sprintf("Anaconda3-%s-MacOSX-%s.sh", latestVersion, arch)
		} else {
			filename = fmt.Sprintf("Anaconda3-%s-MacOSX-%s.sh", latestVersion, arch)
		}
	} else {
		return fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}

	url := fmt.Sprintf("%s/%s", m.scriptURL, filename)
	homeDir, _ := os.UserHomeDir()
	installPath := fmt.Sprintf("%s/.ii/programs/%s", homeDir, program)

	fmt.Printf("下载 Anaconda 安装脚本...\n")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("安装路径: %s\n", installPath)

	// 下载脚本
	tmpFile := "/tmp/anaconda_install.sh"
	var downloadCmd *exec.Cmd

	if utils.CommandExists("wget") {
		downloadCmd = exec.Command("wget", "-O", tmpFile, url)
	} else if utils.CommandExists("curl") {
		downloadCmd = exec.Command("curl", "-L", "-o", tmpFile, url)
	} else {
		return fmt.Errorf("需要 wget 或 curl 来下载安装脚本")
	}

	if output, err := downloadCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("下载失败: %w\n输出: %s", err, output)
	}
	defer os.Remove(tmpFile)

	// 执行安装脚本
	fmt.Printf("\n执行安装脚本...\n")
	installCmd := exec.Command("bash", tmpFile, "-b", "-p", installPath)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("安装失败: %w", err)
	}

	fmt.Printf("\n✓ Anaconda 已安装到: %s\n", installPath)
	fmt.Println("\n请运行以下命令初始化 conda:")
	fmt.Printf("  source %s/etc/profile.d/conda.sh\n", installPath)
	fmt.Println("\n或添加到 ~/.bashrc:")
	fmt.Printf("  echo 'source %s/etc/profile.d/conda.sh' >> ~/.bashrc\n", installPath)

	return nil
}

func (m *ScriptMethod) GetInstallInfo(program, packageName string) string {
	arch := runtime.GOARCH
	if arch == "amd64" {
		arch = "x86_64"
	} else if arch == "arm64" {
		arch = "aarch64"
	}

	homeDir, _ := os.UserHomeDir()
	installPath := fmt.Sprintf("%s/.ii/programs/%s", homeDir, program)
	return fmt.Sprintf("通过官方脚本安装 %s (%s)\n安装位置: %s", program, arch, installPath)
}

func (m *ScriptMethod) Uninstall(ctx context.Context, program, packageName string) error {
	homeDir, _ := os.UserHomeDir()
	installPath := fmt.Sprintf("%s/.ii/programs/%s", homeDir, program)

	fmt.Printf("卸载 %s...\n", program)
	fmt.Printf("安装路径: %s\n", installPath)

	// 检查安装路径是否存在
	if _, err := os.Stat(installPath); os.IsNotExist(err) {
		return fmt.Errorf("%s 未安装在 %s", program, installPath)
	}

	// 删除安装目录
	fmt.Printf("删除安装目录: %s\n", installPath)
	if err := os.RemoveAll(installPath); err != nil {
		return fmt.Errorf("删除失败: %w", err)
	}

	// 提示用户清理配置文件
	fmt.Printf("\n✓ %s 已卸载\n", program)
	fmt.Println("\n请手动清理以下配置文件（如果存在）:")
	fmt.Println("  - ~/.conda")
	fmt.Println("  - ~/.condarc")
	fmt.Println("  - ~/.bashrc 或 ~/.zshrc 中的 conda 初始化代码")

	return nil
}

// GetDownloadURL 获取下载 URL（用于显示信息）
func (m *ScriptMethod) GetDownloadURL() string {
	arch := runtime.GOARCH
	if arch == "amd64" {
		arch = "x86_64"
	} else if arch == "arm64" {
		arch = "aarch64"
	}

	latestVersion := "2024.10-1"
	var filename string

	if runtime.GOOS == "linux" {
		filename = fmt.Sprintf("Anaconda3-%s-Linux-%s.sh", latestVersion, arch)
	} else if runtime.GOOS == "darwin" {
		filename = fmt.Sprintf("Anaconda3-%s-MacOSX-%s.sh", latestVersion, arch)
	}

	return fmt.Sprintf("%s/%s", m.scriptURL, filename)
}

// NewScriptMethod 创建通用脚本安装方法
func NewScriptMethod(name, description string) types.InstallMethod {
	return &ScriptMethod{
		name:        name,
		description: description,
	}
}
