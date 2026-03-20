# ii 开发指南

## 快速开始

### 环境要求

- Go >= 1.21
- Make（可选，用于构建）

### 安装依赖

```bash
make deps
# 或
go mod download
```

### 构建

```bash
make build
```

构建后的二进制文件位于 `bin/ii`

### 运行

```bash
make run
# 或
./bin/ii --help
```

### 测试

```bash
make test
```

## 项目结构

```
ii/
├── cmd/                    # 命令行入口
│   └── ii/                # 主命令
│       └── main.go        # 程序入口
├── internal/              # 内部包（不对外暴露）
│   ├── commands/         # 子命令实现
│   ├── installer/        # 安装器核心逻辑
│   ├── manager/          # 程序管理器
│   ├── config/           # 配置管理
│   └── utils/            # 工具函数
├── pkg/                   # 可对外暴露的包
│   └── types/            # 公共类型定义
├── docs/                  # 文档
├── scripts/               # 构建和安装脚本
├── bin/                   # 构建输出目录
├── go.mod                 # Go 模块定义
├── go.sum                 # 依赖版本锁定
├── Makefile              # 构建脚本
└── README.md             # 项目说明
```

## 开发流程

### 1. 添加新的子命令

在 `internal/commands/` 中创建新的命令文件：

```go
package commands

import (
    "github.com/spf13/cobra"
)

func NewInstallCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "install <program>",
        Short: "安装程序",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            // 实现安装逻辑
        },
    }
    return cmd
}
```

然后在 `cmd/ii/main.go` 中注册：

```go
import "github.com/wii/ii/internal/commands"

func init() {
    rootCmd.AddCommand(commands.NewInstallCmd())
}
```

### 2. 代码风格

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 运行 `make lint` 进行代码检查
- 添加必要的注释

### 3. 提交代码

```bash
# 格式化代码
make fmt

# 代码检查
make lint

# 运行测试
make test

# 提交
git add .
git commit -m "feat: 添加安装命令"
```

## 常用命令

```bash
# 构建
make build

# 运行
make run

# 测试
make test

# 清理
make clean

# 代码检查
make lint

# 格式化
make fmt

# 跨平台构建
make build-all

# 查看帮助
make help
```

## 配置说明

配置文件位于 `~/.ii/config.json`，包含：

- `install_dir`: 程序安装目录
- `sources`: 安装源配置
- `programs`: 已安装程序列表

## 调试

使用 `go run` 直接运行：

```bash
go run ./cmd/ii --help
```

## 发布流程

1. 更新版本号
2. 运行测试
3. 构建所有平台版本：`make build-all`
4. 创建 Git tag
5. 发布到 GitHub Releases

## 注意事项

- 遵循语义化版本规范
- 保持向后兼容
- 充分测试跨平台兼容性
- 文档与代码同步更新
