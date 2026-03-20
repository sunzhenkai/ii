package methods

import (
	"context"
	"fmt"
	"strings"

	"github.com/wii/ii/internal/utils"
	"github.com/wii/ii/pkg/types"
)

// MiseMethod mise 安装方法 (原 rtx)
type MiseMethod struct{}

func NewMiseMethod() types.InstallMethod {
	return &MiseMethod{}
}

func (m *MiseMethod) Name() string {
	return "mise"
}

func (m *MiseMethod) Description() string {
	return "mise (多语言版本管理器)"
}

func (m *MiseMethod) IsAvailable() bool {
	return utils.CommandExists("mise")
}

func (m *MiseMethod) Install(ctx context.Context, program, packageName string) error {
	// mise 使用 mise install 命令
	output, err := utils.RunCommand("mise", "install", packageName)
	if err != nil {
		return fmt.Errorf("安装失败: %w\n输出: %s", err, output)
	}
	return nil
}

func (m *MiseMethod) GetInstallInfo(program, packageName string) string {
	info := fmt.Sprintf("使用 mise 安装 %s (包名: %s)", program, packageName)

	// 获取 mise 数据目录
	if output, err := utils.RunCommand("mise", "data-dir"); err == nil {
		dataDir := strings.TrimSpace(output)
		info += fmt.Sprintf("\n安装位置: %s", dataDir)
	}

	return info
}
