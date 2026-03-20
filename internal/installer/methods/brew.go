package methods

import (
	"context"
	"fmt"

	"github.com/wii/ii/internal/utils"
	"github.com/wii/ii/pkg/types"
)

// BrewMethod Homebrew 安装方法
type BrewMethod struct{}

func NewBrewMethod() types.InstallMethod {
	return &BrewMethod{}
}

func (m *BrewMethod) Name() string {
	return "brew"
}

func (m *BrewMethod) Description() string {
	return "Homebrew (macOS/Linux 包管理器)"
}

func (m *BrewMethod) IsAvailable() bool {
	return utils.CommandExists("brew")
}

func (m *BrewMethod) Install(ctx context.Context, program, packageName string) error {
	output, err := utils.RunCommand("brew", "install", packageName)
	if err != nil {
		return fmt.Errorf("安装失败: %w\n输出: %s", err, output)
	}
	return nil
}

func (m *BrewMethod) GetInstallInfo(program, packageName string) string {
	prefix := ""
	if output, err := utils.RunCommand("brew", "--prefix"); err == nil {
		prefix = output
	}
	return fmt.Sprintf("使用 Homebrew 安装 %s (包名: %s)\n安装位置: %s", program, packageName, prefix)
}
