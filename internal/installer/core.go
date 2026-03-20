package installer

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/wii/ii/internal/installer/methods"
	"github.com/wii/ii/internal/programs"
	"github.com/wii/ii/internal/utils"
	"github.com/wii/ii/pkg/types"
)

// Installer 安装器
type Installer struct {
	registry   *programs.Registry
	osInfo     *utils.OSInfo
	allMethods map[string]types.InstallMethod
}

// NewInstaller 创建安装器
func NewInstaller() *Installer {
	i := &Installer{
		registry:   programs.NewRegistry(),
		osInfo:     utils.GetOSInfo(),
		allMethods: make(map[string]types.InstallMethod),
	}

	// 注册所有安装方法
	i.registerMethods()

	return i
}

// registerMethods 注册所有安装方法
func (i *Installer) registerMethods() {
	// 获取系统包管理器
	pm := utils.GetPackageManager()
	if pm != "" && pm != "brew" {
		// 系统包管理器（非 brew）
		i.allMethods[pm] = methods.NewPackageManagerMethod(pm)
	}

	// 注册其他安装方法
	i.allMethods["brew"] = methods.NewBrewMethod()
	i.allMethods["mise"] = methods.NewMiseMethod()
	i.allMethods["asdf"] = methods.NewAsdfMethod()
}

// GetAvailableMethods 获取程序可用的安装方法
func (i *Installer) GetAvailableMethods(program types.Program) []types.InstallMethod {
	supportedMethods := program.GetInstallMethods()
	available := []types.InstallMethod{}

	// 按优先级排序：系统包管理器 > brew > mise > asdf
	priority := []string{
		utils.GetPackageManager(), // 系统包管理器优先
		"brew",
		"mise",
		"asdf",
	}

	for _, methodName := range priority {
		if packageName, ok := supportedMethods[methodName]; ok {
			if method, exists := i.allMethods[methodName]; exists {
				// 检查该方法是否可用
				if method.IsAvailable() {
					// 临时保存包名信息
					available = append(available, &methodWithPackage{
						InstallMethod: method,
						packageName:   packageName,
					})
				}
			}
		}
	}

	return available
}

// methodWithPackage 带包名信息的安装方法
type methodWithPackage struct {
	types.InstallMethod
	packageName string
}

func (m *methodWithPackage) PackageName() string {
	return m.packageName
}

// InstallProgram 安装程序
func (i *Installer) InstallProgram(ctx context.Context, programName string, opts types.InstallOption) error {
	// 获取程序信息
	program, err := i.registry.Get(programName)
	if err != nil {
		return err
	}

	// 检查平台支持
	if !i.isPlatformSupported(program) {
		return fmt.Errorf("程序 %s 不支持当前平台 (%s/%s)",
			programName, i.osInfo.Type, i.osInfo.Arch)
	}

	// 获取可用的安装方法
	availableMethods := i.GetAvailableMethods(program)
	if len(availableMethods) == 0 {
		return fmt.Errorf("没有找到可用的安装方法")
	}

	// 如果指定了安装方法
	if opts.Method != "" {
		return i.installWithMethod(ctx, program, opts.Method, opts)
	}

	// 显示可用方法并让用户选择
	selectedMethod, err := i.selectMethod(program, availableMethods, opts)
	if err != nil {
		return err
	}

	// 执行安装
	return i.doInstall(ctx, program, selectedMethod, opts)
}

// installWithMethod 使用指定的方法安装
func (i *Installer) installWithMethod(ctx context.Context, program types.Program, methodName string, opts types.InstallOption) error {
	supportedMethods := program.GetInstallMethods()
	packageName, ok := supportedMethods[methodName]
	if !ok {
		return fmt.Errorf("程序 %s 不支持使用 %s 安装", program.Name(), methodName)
	}

	method, exists := i.allMethods[methodName]
	if !exists {
		return fmt.Errorf("未知的安装方法: %s", methodName)
	}

	if !method.IsAvailable() {
		return fmt.Errorf("安装方法 %s 在当前系统上不可用", methodName)
	}

	return i.doInstall(ctx, program, &methodWithPackage{
		InstallMethod: method,
		packageName:   packageName,
	}, opts)
}

// selectMethod 选择安装方法
func (i *Installer) selectMethod(program types.Program, availableMethods []types.InstallMethod, opts types.InstallOption) (types.InstallMethod, error) {
	// 如果只有一个方法且指定了自动确认
	if len(availableMethods) == 1 && opts.Yes {
		return availableMethods[0], nil
	}

	// 显示可用方法
	fmt.Printf("\n程序: %s\n", program.Name())
	fmt.Printf("描述: %s\n\n", program.Description())
	fmt.Println("可用的安装方法:")
	fmt.Println(strings.Repeat("-", 60))

	supportedMethods := program.GetInstallMethods()

	for idx, method := range availableMethods {
		// 获取包名
		var packageName string
		if m, ok := method.(*methodWithPackage); ok {
			packageName = m.packageName
		} else {
			packageName = supportedMethods[method.Name()]
		}

		fmt.Printf("\n%d. %s\n", idx+1, method.Description())
		fmt.Printf("   %s\n", method.GetInstallInfo(program.Name(), packageName))
	}

	// 如果指定了自动确认，选择第一个
	if opts.Yes {
		fmt.Printf("\n自动选择: %s\n", availableMethods[0].Name())
		return availableMethods[0], nil
	}

	// 如果是 dry-run 模式，直接返回第一个
	if opts.DryRun {
		fmt.Printf("\n[Dry Run] 将使用: %s\n", availableMethods[0].Name())
		return availableMethods[0], nil
	}

	// 让用户选择
	fmt.Printf("\n请选择安装方法 [1-%d] (默认: 1): ", len(availableMethods))

	var choice int
	choice = 1 // 默认选择第一个

	if _, err := fmt.Scanf("%d", &choice); err != nil {
		// 用户直接回车，使用默认值
		choice = 1
	}

	if choice < 1 || choice > len(availableMethods) {
		return nil, fmt.Errorf("无效的选择: %d", choice)
	}

	return availableMethods[choice-1], nil
}

// doInstall 执行安装
func (i *Installer) doInstall(ctx context.Context, program types.Program, method types.InstallMethod, opts types.InstallOption) error {
	// 获取包名
	var packageName string
	if m, ok := method.(*methodWithPackage); ok {
		packageName = m.packageName
	} else {
		supportedMethods := program.GetInstallMethods()
		packageName = supportedMethods[method.Name()]
	}

	fmt.Printf("\n开始安装...\n")
	fmt.Printf("程序: %s\n", program.Name())
	fmt.Printf("方法: %s\n", method.Description())
	fmt.Printf("包名: %s\n\n", packageName)

	// 如果是 dry-run 模式，不实际执行
	if opts.DryRun {
		fmt.Println("[Dry Run] 跳过实际安装")
		return nil
	}

	// 确认安装（如果没有自动确认）
	if !opts.Yes {
		fmt.Printf("确认安装? [y/N]: ")
		var confirm string
		fmt.Scanf("%s", &confirm)
		if strings.ToLower(confirm) != "y" {
			return fmt.Errorf("用户取消安装")
		}
	}

	// 执行安装
	err := method.Install(ctx, program.Name(), packageName)
	if err != nil {
		return fmt.Errorf("安装失败: %w", err)
	}

	fmt.Printf("\n✓ %s 安装成功!\n", program.Name())

	// 显示使用示例
	fmt.Printf("\n使用示例:\n")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println(program.GetUsage())

	return nil
}

// isPlatformSupported 检查平台是否支持
func (i *Installer) isPlatformSupported(program types.Program) bool {
	supported := program.GetSupportedPlatforms()

	archs, ok := supported[i.osInfo.Type]
	if !ok {
		return false
	}

	for _, arch := range archs {
		if arch == i.osInfo.Arch || arch == "*" {
			return true
		}
	}

	return false
}

// ListPrograms 列出所有支持的程序
func (i *Installer) ListPrograms() {
	programs := i.registry.List()

	fmt.Println("支持的程序:")
	fmt.Println(strings.Repeat("-", 60))

	// 按名称排序
	sort.Slice(programs, func(i, j int) bool {
		return programs[i].Name() < programs[j].Name()
	})

	for _, p := range programs {
		fmt.Printf("  %-20s %s\n", p.Name(), p.Description())
	}
}

// GetProgram 获取程序信息
func (i *Installer) GetProgram(name string) (types.Program, error) {
	return i.registry.Get(name)
}
