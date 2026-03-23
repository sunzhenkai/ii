package programs

import (
	"runtime"

	"github.com/wii/ii/pkg/types"
)

// Ollama 程序定义
type Ollama struct{}

func NewOllama() types.Program {
	return &Ollama{}
}

func (o *Ollama) Name() string {
	return "ollama"
}

func (o *Ollama) Description() string {
	return "Ollama 是一个用于在本地运行大语言模型的工具，支持 Llama 2、Mistral、Gemma 等模型"
}

func (o *Ollama) GetInstallMethods() map[string]string {
	// 返回支持的安装方法和对应的包名
	// 格式: map[安装方法名]包名
	methods := map[string]string{
		// Linux/macOS - 官方脚本安装（推荐）
		"ollama-script": "ollama",
	}

	// macOS 可以使用 Homebrew
	if runtime.GOOS == "darwin" {
		methods["brew"] = "ollama"
	}

	// Linux 也可以使用 Homebrew
	if runtime.GOOS == "linux" {
		methods["brew"] = "ollama"
	}

	return methods
}

func (o *Ollama) GetSupportedPlatforms() map[string][]string {
	return map[string][]string{
		"linux":  {"amd64", "arm64"},
		"darwin": {"amd64", "arm64"},
	}
}

func (o *Ollama) GetUsage() string {
	return `常用命令:

1. 启动 Ollama 服务:
   ollama serve

2. 拉取模型:
   ollama pull llama2
   ollama pull mistral
   ollama pull gemma
   ollama pull codellama

3. 运行模型（交互式）:
   ollama run llama2
   ollama run mistral

4. 列出已下载的模型:
   ollama list

5. 删除模型:
   ollama rm llama2

6. 显示模型信息:
   ollama show llama2

7. 从 Modelfile 创建模型:
   ollama create mymodel -f Modelfile

8. 推送模型到仓库:
   ollama push mymodel

9. 拉取模型到本地:
   ollama pull mymodel

10. 复制模型:
    ollama cp llama2 my-llama2

11. API 使用（默认端口 11434）:
    curl http://localhost:11434/api/generate -d '{
      "model": "llama2",
      "prompt": "Why is the sky blue?"
    }'

12. 查看运行状态:
    ps aux | grep ollama

常用模型:
    llama2          - Meta 的 Llama 2 模型
    mistral         - Mistral AI 的模型
    gemma           - Google 的 Gemma 模型
    codellama       - Code Llama，专用于代码
    llama2-uncensored - 无审查版本
    orca-mini       - Microsoft 的 Orca 模型

更多信息:
    ollama --help
    https://ollama.com/
    https://github.com/ollama/ollama
    https://ollama.com/library`
}
