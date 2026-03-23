package programs

import (
	"runtime"

	"github.com/wii/ii/pkg/types"
)

// Anaconda3 程序定义
type Anaconda3 struct{}

func NewAnaconda3() types.Program {
	return &Anaconda3{}
}

func (a *Anaconda3) Name() string {
	return "anaconda3"
}

func (a *Anaconda3) Description() string {
	return "Anaconda 是一个用于科学计算的 Python 发行版，支持 Linux, macOS, Windows"
}

func (a *Anaconda3) GetInstallMethods() map[string]string {
	// 返回支持的安装方法和对应的包名
	// 格式: map[安装方法名]包名

	// 根据操作系统返回不同的安装方法
	methods := map[string]string{
		// Linux/macOS - 官方脚本安装（推荐）
		"anaconda-script": "anaconda3",
	}

	// macOS 可以额外使用 Homebrew Cask
	// 注意：Homebrew 的 anaconda cask 只支持 macOS
	if runtime.GOOS == "darwin" {
		methods["brew"] = "anaconda"
	}

	return methods
}

func (a *Anaconda3) GetSupportedPlatforms() map[string][]string {
	return map[string][]string{
		"linux":  {"amd64", "arm64"},
		"darwin": {"amd64", "arm64"},
	}
}

func (a *Anaconda3) GetUsage() string {
	return `常用命令:

1. 初始化 conda（首次安装后）:
   source ~/.ii/programs/anaconda3/etc/profile.d/conda.sh
   # 或添加到 ~/.bashrc:
   echo 'source ~/.ii/programs/anaconda3/etc/profile.d/conda.sh' >> ~/.bashrc

2. 创建新的虚拟环境:
   conda create -n myenv python=3.10
   conda create -n myenv numpy pandas matplotlib

3. 激活环境:
   conda activate myenv

4. 退出环境:
   conda deactivate

5. 列出所有环境:
   conda env list
   conda info --envs

6. 删除环境:
   conda remove -n myenv --all

7. 安装包:
   conda install numpy
   conda install -c conda-forge pandas

8. 搜索包:
   conda search numpy

9. 更新 conda:
   conda update conda
   conda update anaconda

10. 配置国内镜像源（加速下载）:
    conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main
    conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free
    conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/conda-forge
    conda config --set show_channel_urls yes

11. pip 配置镜像源:
    pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

更多信息:
    conda --help
    https://docs.anaconda.com/
    https://conda.io/projects/conda/en/latest/user-guide/index.html`
}
