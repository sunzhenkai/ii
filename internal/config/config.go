package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// DefaultConfig 默认配置
const (
	DefaultInstallDir = "~/.ii/programs"
	ConfigFileName    = "config.json"
	InstalledFileName = "installed.json"
)

// InstalledProgram 已安装程序记录
type InstalledProgram struct {
	Name        string    `json:"name"`         // 程序名称
	PackageName string    `json:"package_name"` // 包名
	Method      string    `json:"method"`       // 安装方法
	Version     string    `json:"version"`      // 版本号（如果有）
	InstallTime time.Time `json:"install_time"` // 安装时间
	InstallPath string    `json:"install_path"` // 安装路径（如果有）
}

// InstalledDB 已安装程序数据库
type InstalledDB struct {
	Programs map[string]InstalledProgram `json:"programs"`
}

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

// GetInstalledDBPath 获取已安装程序数据库路径
func (cm *ConfigManager) GetInstalledDBPath() string {
	return filepath.Join(cm.configDir, InstalledFileName)
}

// LoadInstalledDB 加载已安装程序数据库
func (cm *ConfigManager) LoadInstalledDB() (*InstalledDB, error) {
	dbPath := cm.GetInstalledDBPath()

	// 如果文件不存在，返回空数据库
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return &InstalledDB{Programs: make(map[string]InstalledProgram)}, nil
	}

	data, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}

	var db InstalledDB
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}

	// 确保 Programs 不为 nil
	if db.Programs == nil {
		db.Programs = make(map[string]InstalledProgram)
	}

	return &db, nil
}

// SaveInstalledDB 保存已安装程序数据库
func (cm *ConfigManager) SaveInstalledDB(db *InstalledDB) error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cm.GetInstalledDBPath(), data, 0644)
}

// RecordInstall 记录安装信息
func (cm *ConfigManager) RecordInstall(program, packageName, method string) error {
	db, err := cm.LoadInstalledDB()
	if err != nil {
		return err
	}

	db.Programs[program] = InstalledProgram{
		Name:        program,
		PackageName: packageName,
		Method:      method,
		InstallTime: time.Now(),
	}

	return cm.SaveInstalledDB(db)
}

// GetInstalledProgram 获取已安装程序信息
func (cm *ConfigManager) GetInstalledProgram(program string) (*InstalledProgram, error) {
	db, err := cm.LoadInstalledDB()
	if err != nil {
		return nil, err
	}

	info, exists := db.Programs[program]
	if !exists {
		return nil, nil
	}

	return &info, nil
}

// RemoveInstalledProgram 移除已安装程序记录
func (cm *ConfigManager) RemoveInstalledProgram(program string) error {
	db, err := cm.LoadInstalledDB()
	if err != nil {
		return err
	}

	delete(db.Programs, program)

	return cm.SaveInstalledDB(db)
}

// ListInstalledPrograms 列出所有已安装程序
func (cm *ConfigManager) ListInstalledPrograms() ([]InstalledProgram, error) {
	db, err := cm.LoadInstalledDB()
	if err != nil {
		return nil, err
	}

	programs := make([]InstalledProgram, 0, len(db.Programs))
	for _, p := range db.Programs {
		programs = append(programs, p)
	}

	return programs, nil
}
