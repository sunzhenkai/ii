package config

import (
	"os"
	"path/filepath"
)

// DefaultConfig 默认配置
const (
	DefaultInstallDir = "~/.ii/programs"
	ConfigFileName    = "config.json"
)

// ConfigManager 配置管理器
type ConfigManager struct {
	configPath string
	configDir  string
}

// NewConfigManager 创建配置管理器
func NewConfigManager() *ConfigManager {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".ii")

	return &ConfigManager{
		configPath: filepath.Join(configDir, ConfigFileName),
		configDir:  configDir,
	}
}

// Init 初始化配置目录
func (cm *ConfigManager) Init() error {
	// 创建配置目录
	if err := os.MkdirAll(cm.configDir, 0755); err != nil {
		return err
	}

	// 创建安装目录
	installDir := cm.GetInstallDir()
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return err
	}

	return nil
}

// GetInstallDir 获取安装目录
func (cm *ConfigManager) GetInstallDir() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".ii", "programs")
}

// GetConfigPath 获取配置文件路径
func (cm *ConfigManager) GetConfigPath() string {
	return cm.configPath
}

// GetConfigDir 获取配置目录
func (cm *ConfigManager) GetConfigDir() string {
	return cm.configDir
}
