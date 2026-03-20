package types

// Program 表示一个可管理的程序
type Program struct {
	Name        string            `json:"name"`         // 程序名称
	Version     string            `json:"version"`      // 版本号
	Description string            `json:"description"`  // 描述信息
	InstallPath string            `json:"install_path"` // 安装路径
	BinaryName  string            `json:"binary_name"`  // 二进制文件名
	Source      string            `json:"source"`       // 安装源
	Metadata    map[string]string `json:"metadata"`     // 额外元数据
}

// InstallOptions 安装选项
type InstallOptions struct {
	Version     string // 指定版本
	Source      string // 安装源
	Force       bool   // 强制重新安装
	SkipVerify  bool   // 跳过校验
	InstallPath string // 自定义安装路径
}

// Config 配置结构
type Config struct {
	InstallDir string            `json:"install_dir"` // 安装目录
	Sources    map[string]string `json:"sources"`     // 源地址映射
	Programs   []Program         `json:"programs"`    // 已安装程序列表
}
