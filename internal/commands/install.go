package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wii/ii/internal/installer"
	"github.com/wii/ii/pkg/types"
)

var (
	installMethod string
	installYes    bool
	installDryRun bool
)

// NewInstallCmd 创建 install 命令
func NewInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install <program>",
		Short: "安装程序",
		Long:  "使用最佳的安装方式安装指定程序",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			programName := args[0]

			// 创建安装器
			inst := installer.NewInstaller()

			// 安装选项
			opts := types.InstallOption{
				Method: installMethod,
				Yes:    installYes,
				DryRun: installDryRun,
			}

			// 执行安装
			return inst.InstallProgram(context.Background(), programName, opts)
		},
	}

	// 命令行参数
	cmd.Flags().StringVarP(&installMethod, "method", "m", "", "指定安装方法 (apt/brew/mise/asdf)")
	cmd.Flags().BoolVarP(&installYes, "yes", "y", false, "自动确认，无需交互")
	cmd.Flags().BoolVarP(&installDryRun, "dry-run", "d", false, "只展示将要执行的操作，不实际执行")

	return cmd
}

// NewListCmd 创建 list 命令
func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "列出支持的程序",
		Long:  "列出所有可以通过 ii 安装的程序",
		Run: func(cmd *cobra.Command, args []string) {
			inst := installer.NewInstaller()
			inst.ListPrograms()
		},
	}

	return cmd
}

// NewSearchCmd 创建 search 命令
func NewSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <keyword>",
		Short: "搜索程序",
		Long:  "搜索支持的程序",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			keyword := args[0]
			fmt.Printf("搜索程序: %s\n", keyword)
			// TODO: 实现搜索逻辑
			return fmt.Errorf("搜索功能尚未实现")
		},
	}

	return cmd
}
