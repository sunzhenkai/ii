# ii

ii 是一个跨系统的程序管理工具，提供统一的程序安装功能，以及更简单快捷的程序管理体验。

## 特性

- 🌍 **跨平台支持** - 支持 Linux、macOS、Windows
- 📦 **统一安装** - 一键安装各种开发工具和程序
- 🔄 **版本管理** - 轻松管理和切换程序版本
- 🚀 **简单易用** - 直观的命令行界面

## 快速开始

### 安装

从 [Releases](https://github.com/wii/ii/releases) 页面下载对应平台的二进制文件。

### 从源码构建

```bash
# 克隆仓库
git clone https://github.com/wii/ii.git
cd ii

# 安装依赖
make deps

# 构建
make build

# 运行
./bin/ii --help
```

## 使用

```bash
# 查看帮助
ii --help

# 安装程序（待实现）
ii install <program>

# 列出已安装程序（待实现）
ii list

# 更新程序（待实现）
ii update <program>
```

## 开发

详细开发指南请参考 [开发文档](docs/development.md)。

```bash
# 运行测试
make test

# 代码检查
make lint

# 跨平台构建
make build-all
```

## 技术栈

- **Go** - 高性能、跨平台的编程语言
- **Cobra** - 强大的命令行框架

更多技术细节请参考 [技术架构文档](docs/tech-stack.md)。

## 许可证

[MIT License](LICENSE)

## 贡献

欢迎提交 Issue 和 Pull Request！
