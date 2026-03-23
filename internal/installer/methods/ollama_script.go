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

// OllamaScriptMethod Ollama 脚本安装方法
type OllamaScriptMethod struct{}

// NewOllamaScriptMethod 创建 Ollama 脚本安装方法
func NewOllamaScriptMethod() types.InstallMethod {
	return &OllamaScriptMethod{}
}

func (m *OllamaScriptMethod) Name() string {
	return "ollama-script"
}

func (m *OllamaScriptMethod) Description() string {
	return "官方安装脚本"
}

func (m *OllamaScriptMethod) IsAvailable() bool {
	// 需要 curl 或 wget
	return utils.CommandExists("curl") || utils.CommandExists("wget")
}

func (m *OllamaScriptMethod) Install(ctx context.Context, program, packageName string) error {
	homeDir, _ := os.UserHomeDir()
	installPath := fmt.Sprintf("%s/.ii/programs/%s", homeDir, program)

	fmt.Printf("安装 Ollama...\n")
	fmt.Printf("安装路径: %s\n", installPath)

	// Ollama 官方安装脚本
	scriptURL := "https://ollama.com/install.sh"

	// 下载并执行安装脚本
	var cmd *exec.Cmd
	if utils.CommandExists("curl") {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("curl -fsSL %s | sh", scriptURL))
	} else if utils.CommandExists("wget") {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("wget -qO- %s | sh", scriptURL))
	} else {
		return fmt.Errorf("需要 curl 或 wget 来下载安装脚本")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\n下载并执行官方安装脚本...\n")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("安装失败: %w", err)
	}

	fmt.Printf("\n✓ Ollama 安装成功!\n")
	fmt.Println("\n常用命令:")
	fmt.Println("  ollama serve      - 启动服务")
	fmt.Println("  ollama run llama2 - 运行模型")
	fmt.Println("  ollama list       - 列出模型")
	fmt.Println("\n更多信息:")
	fmt.Println("  https://ollama.com/")
	fmt.Println("  https://github.com/ollama/ollama")

	return nil
}

func (m *OllamaScriptMethod) GetInstallInfo(program, packageName string) string {
	arch := runtime.GOARCH
	if arch == "amd64" {
		arch = "x86_64"
	} else if arch == "arm64" {
		arch = "aarch64"
	}

	return fmt.Sprintf("通过官方脚本安装 %s (%s)\n安装位置: 系统路径 (/usr/local/bin)", program, arch)
}

func (m *OllamaScriptMethod) Uninstall(ctx context.Context, program, packageName string) error {
	fmt.Printf("卸载 Ollama...\n")

	// 停止 ollama 服务
	fmt.Println("停止 ollama 服务...")
	if utils.CommandExists("systemctl") {
		exec.Command("systemctl", "stop", "ollama").Run()
		exec.Command("systemctl", "disable", "ollama").Run()
	}

	// 删除 ollama 二进制文件
	ollamaPath := "/usr/local/bin/ollama"
	if _, err := os.Stat(ollamaPath); err == nil {
		fmt.Printf("删除二进制文件: %s\n", ollamaPath)
		if err := exec.Command("sudo", "rm", "-f", ollamaPath).Run(); err != nil {
			fmt.Printf("警告: 无法删除 %s: %v\n", ollamaPath, err)
		}
	}

	// 删除服务文件
	serviceFiles := []string{
		"/etc/systemd/system/ollama.service",
		"/usr/lib/systemd/system/ollama.service",
	}
	for _, serviceFile := range serviceFiles {
		if _, err := os.Stat(serviceFile); err == nil {
			fmt.Printf("删除服务文件: %s\n", serviceFile)
			exec.Command("sudo", "rm", "-f", serviceFile).Run()
		}
	}

	// 重新加载 systemd
	if utils.CommandExists("systemctl") {
		exec.Command("systemctl", "daemon-reload").Run()
	}

	// 提示用户清理模型和数据
	fmt.Println("\n✓ Ollama 已卸载")
	fmt.Println("\n请手动清理以下目录（如果需要）:")
	fmt.Println("  - ~/.ollama        (模型和配置)")
	fmt.Println("  - /usr/share/ollama (共享数据)")

	return nil
}
