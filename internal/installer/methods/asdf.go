package methods

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wii/ii/internal/utils"
	"github.com/wii/ii/pkg/types"
)

// AsdfMethod asdf 安装方法
type AsdfMethod struct{}

func NewAsdfMethod() types.InstallMethod {
	return &AsdfMethod{}
}

func (m *AsdfMethod) Name() string {
	return "asdf"
}

func (m *AsdfMethod) Description() string {
	return "asdf (多语言版本管理器)"
}

func (m *AsdfMethod) IsAvailable() bool {
	return utils.CommandExists("asdf")
}

func (m *AsdfMethod) Install(ctx context.Context, program, packageName string) error {
	// asdf 需要先添加插件，然后安装
	// packageName 格式可能是 "plugin" 或 "plugin/version"
	parts := strings.Split(packageName, "/")
	pluginName := parts[0]

	// 检查插件是否已安装
	output, err := utils.RunCommand("asdf", "plugin", "list")
	if err != nil {
		return fmt.Errorf("检查插件列表失败: %w", err)
	}

	// 如果插件未安装，先添加插件
	if !strings.Contains(output, pluginName) {
		output, err = utils.RunCommand("asdf", "plugin", "add", pluginName)
		if err != nil {
			return fmt.Errorf("添加插件失败: %w\n输出: %s", err, output)
		}
	}

	// 安装包
	installArgs := []string{"install"}
	installArgs = append(installArgs, parts...)

	output, err = utils.RunCommand("asdf", installArgs...)
	if err != nil {
		return fmt.Errorf("安装失败: %w\n输出: %s", err, output)
	}

	return nil
}

func (m *AsdfMethod) GetInstallInfo(program, packageName string) string {
	info := fmt.Sprintf("使用 asdf 安装 %s (包名: %s)", program, packageName)

	// 获取 asdf 安装目录
	homeDir, _ := os.UserHomeDir()
	asdfDir := filepath.Join(homeDir, ".asdf")
	if dir := os.Getenv("ASDF_DATA_DIR"); dir != "" {
		asdfDir = dir
	}

	info += fmt.Sprintf("\n安装位置: %s", asdfDir)
	return info
}
