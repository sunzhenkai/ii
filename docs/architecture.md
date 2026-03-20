# ii 架构设计

## 设计原则

### 1. 接口驱动设计

所有核心功能都通过接口定义，便于扩展和测试：

- `types.InstallMethod` - 安装方法接口
- `types.Program` - 程序定义接口

### 2. 分层架构

```
cmd/ii/                 # 命令行入口层
  └── main.go
internal/
  ├── commands/         # 命令处理层
  ├── installer/        # 安装器核心层
  │   └── methods/     # 安装方法实现
  ├── programs/        # 程序定义层
  └── utils/           # 工具层
pkg/types/              # 公共类型定义
```

### 3. 开闭原则

对扩展开放，对修改关闭：

- 添加新的安装方法：只需实现 `InstallMethod` 接口
- 添加新的程序：只需实现 `Program` 接口并在注册表中注册
- 无需修改核心代码

## 核心组件

### 1. 安装方法 (InstallMethod)

位置：`pkg/types/types.go`

```go
type InstallMethod interface {
    Name() string
    Description() string
    IsAvailable() bool
    Install(ctx context.Context, program, packageName string) error
    GetInstallInfo(program, packageName string) string
}
```

已实现的安装方法：
- **PackageManagerMethod** - 系统包管理器（apt/yum/dnf/pacman/zypper）
- **BrewMethod** - Homebrew
- **MiseMethod** - mise
- **AsdfMethod** - asdf

### 2. 程序定义 (Program)

位置：`pkg/types/types.go`

```go
type Program interface {
    Name() string
    Description() string
    GetInstallMethods() map[string]string
    GetSupportedPlatforms() map[string][]string
}
```

已实现的程序：
- **WireGuard** - VPN 工具

### 3. 安装器核心 (Installer)

位置：`internal/installer/core.go`

职责：
- 管理所有安装方法
- 管理程序注册表
- 自动选择最佳安装方法
- 用户交互和确认
- 执行安装

### 4. 程序注册表 (Registry)

位置：`internal/programs/registry.go`

职责：
- 注册所有支持的程序
- 提供程序查询功能

## 添加新程序

1. 在 `internal/programs/` 创建新文件，例如 `nodejs.go`：

```go
package programs

import "github.com/wii/ii/pkg/types"

type NodeJS struct{}

func NewNodeJS() types.Program {
    return &NodeJS{}
}

func (n *NodeJS) Name() string {
    return "nodejs"
}

func (n *NodeJS) Description() string {
    return "Node.js JavaScript 运行时"
}

func (n *NodeJS) GetInstallMethods() map[string]string {
    return map[string]string{
        "apt":    "nodejs",
        "brew":   "node",
        "mise":   "node",
        "asdf":   "nodejs",
    }
}

func (n *NodeJS) GetSupportedPlatforms() map[string][]string {
    return map[string][]string{
        "linux":  {"amd64", "arm64"},
        "darwin": {"amd64", "arm64"},
    }
}

func (n *NodeJS) GetUsage() string {
    return `常用命令:

1. 查看版本:
   node --version
   npm --version

2. 运行 JavaScript 文件:
   node app.js

3. 初始化项目:
   npm init
   npm init -y

4. 安装依赖:
   npm install
   npm install <package>
   npm install -D <package>

5. 运行脚本:
   npm run <script>

更多信息:
   node --help
   npm help
   https://nodejs.org/`
}
```

2. 在 `internal/programs/registry.go` 中注册：

```go
func (r *Registry) registerAll() {
    r.Register(NewWireGuard())
    r.Register(NewNodeJS())  // 添加这行
}
```

## 添加新安装方法

1. 在 `internal/installer/methods/` 创建新文件：

```go
package methods

import (
    "context"
    "fmt"
    "github.com/wii/ii/pkg/types"
)

type MyMethod struct{}

func NewMyMethod() types.InstallMethod {
    return &MyMethod{}
}

func (m *MyMethod) Name() string {
    return "my-method"
}

func (m *MyMethod) Description() string {
    return "My Custom Method"
}

func (m *MyMethod) IsAvailable() bool {
    // 检查方法是否可用
    return true
}

func (m *MyMethod) Install(ctx context.Context, program, packageName string) error {
    // 实现安装逻辑
    return nil
}

func (m *MyMethod) GetInstallInfo(program, packageName string) string {
    return fmt.Sprintf("使用 MyMethod 安装 %s", program)
}
```

2. 在 `internal/installer/core.go` 中注册：

```go
func (i *Installer) registerMethods() {
    // ... 其他方法
    i.allMethods["my-method"] = methods.NewMyMethod()
}
```

## 代码复用

### 安装方法复用

所有安装方法实现相同的接口，可以在不同程序间复用。

### 程序定义复用

程序定义独立于安装方法，可以支持多种安装方式。

### 工具函数复用

`internal/utils/` 提供系统检测、命令执行等通用功能。

## 错误处理

- 清晰的错误信息，便于理解和定位问题
- 错误向上传递，由命令层统一处理
- 使用 `fmt.Errorf` 包装错误上下文

## 扩展性

### 支持版本管理

可以为 `Program` 接口添加版本相关方法：

```go
type Program interface {
    // ... 现有方法
    GetAvailableVersions() []string
    GetLatestVersion() string
}
```

### 支持配置文件

可以为程序添加配置支持：

```go
type Program interface {
    // ... 现有方法
    GetConfig() *ProgramConfig
}
```

### 支持插件系统

可以设计插件接口，动态加载程序定义：

```go
type ProgramPlugin interface {
    Init() error
    GetPrograms() []Program
}
```

## 测试策略

### 单元测试

- 每个安装方法独立测试
- 程序定义测试
- 工具函数测试

### 集成测试

- 安装流程测试
- 用户交互测试（使用 mock）

### 端到端测试

- 真实环境安装测试
- 跨平台测试
