package types

import "context"

// InstallMethod 安装方法接口
// 所有安装方法（包管理器、brew、mise等）都需要实现这个接口
type InstallMethod interface {
	// Name 返回安装方法名称（如 "apt", "brew", "mise"）
	Name() string

	// Description 返回安装方法描述
	Description() string

	// IsAvailable 检查该方法在当前系统上是否可用
	IsAvailable() bool

	// Install 执行安装
	// program: 程序名称
	// packageName: 在该安装方法中的包名（可能与程序名不同）
	Install(ctx context.Context, program, packageName string) error

	// GetInstallInfo 获取安装信息（用于向用户展示）
	// 返回：安装位置、版本等信息的描述
	GetInstallInfo(program, packageName string) string
}

// Program 程序定义接口
// 所有可安装程序都需要实现这个接口
type Program interface {
	// Name 返回程序名称
	Name() string

	// Description 返回程序描述
	Description() string

	// GetInstallMethods 返回支持的安装方法
	// 返回: map[安装方法名]包名
	// 例如: {"apt": "wireguard", "brew": "wireguard-tools"}
	GetInstallMethods() map[string]string

	// GetSupportedPlatforms 返回支持的平台
	// 返回: map[操作系统][]架构，如 {"linux": ["amd64", "arm64"]}
	GetSupportedPlatforms() map[string][]string
}

// ProgramInfo 程序信息（用于存储和展示）
type ProgramInfo struct {
	Name        string            `json:"name"`         // 程序名称
	Version     string            `json:"version"`      // 版本号
	Description string            `json:"description"`  // 描述信息
	InstallPath string            `json:"install_path"` // 安装路径
	BinaryName  string            `json:"binary_name"`  // 二进制文件名
	Source      string            `json:"source"`       // 安装源
	Metadata    map[string]string `json:"metadata"`     // 额外元数据
}

// InstallOption 安装选项
type InstallOption struct {
	Method string // 指定安装方法，为空则自动选择
	Force  bool   // 强制重新安装
	Yes    bool   // 自动确认，不需要用户交互
	DryRun bool   // 只展示将要执行的操作，不实际执行
}

// InstallResult 安装结果
type InstallResult struct {
	Success bool
	Method  string
	Message string
	Error   error
}

// Config 配置结构
type Config struct {
	InstallDir string            `json:"install_dir"` // 安装目录
	Sources    map[string]string `json:"sources"`     // 源地址映射
	Programs   []ProgramInfo     `json:"programs"`    // 已安装程序列表
}
