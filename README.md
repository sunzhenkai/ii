# ii

ii 是一个跨系统的程序管理工具，提供统一的程序安装功能，以及更简单快捷的程序管理体验。

## 特性

- 🌍 **跨平台支持** - 支持 Linux、macOS、Windows
- 📦 **统一安装** - 一键安装各种开发工具和程序
- 🔄 **版本管理** - 轻松管理和切换程序版本
- 🚀 **简单易用** - 直观的命令行界面

## 快速开始

### 安装

#### 方式一：从源码安装（推荐）

```bash
# 克隆仓库
git clone https://github.com/wii/ii.git
cd ii

# 安装依赖
make deps

# 安装到 ~/.local/bin
make install

# 确保 ~/.local/bin 在 PATH 中
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# 验证安装
ii --version
```

#### 方式二：下载预编译二进制

从 [Releases](https://github.com/wii/ii/releases) 页面下载对应平台的二进制文件，然后：

```bash
chmod +x ii
mv ii ~/.local/bin/
```

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

# 或者直接安装
make install
```

## 使用

```bash
# 查看帮助
ii --help

# 列出支持的程序
ii list

# 安装程序（交互式选择安装方法）
ii install wireguard

# 安装程序（自动确认，使用最佳安装方法）
ii install wireguard -y

# 安装程序（指定安装方法）
ii install wireguard --method brew

# 安装程序（dry-run 模式，只查看将要执行的操作）
ii install wireguard --dry-run

# 查看安装命令帮助
ii install --help
```

### 支持的安装方法

ii 会自动检测系统上可用的安装方法，并按以下优先级选择：

1. **系统包管理器** - apt, yum, dnf, pacman, zypper 等
2. **Homebrew** - macOS/Linux 包管理器
3. **mise** - 多语言版本管理器
4. **asdf** - 多语言版本管理器

### 已支持的程序

- **wireguard** - 快速、现代的 VPN 技术

更多程序持续添加中...

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
