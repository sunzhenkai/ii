package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// 版本信息，构建时通过 ldflags 注入
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "ii",
	Short: "跨系统的程序管理工具",
	Long: `ii 提供跨系统统一的程序安装功能，
以及以更简单快捷的方式使用程序常用功能。`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		// 如果没有指定子命令，显示帮助信息
		cmd.Help()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
