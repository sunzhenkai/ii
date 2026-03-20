package programs

import (
	"fmt"

	"github.com/wii/ii/pkg/types"
)

// Registry 程序注册表
type Registry struct {
	programs map[string]types.Program
}

// NewRegistry 创建程序注册表
func NewRegistry() *Registry {
	r := &Registry{
		programs: make(map[string]types.Program),
	}

	// 注册所有支持的程序
	r.registerAll()

	return r
}

// Register 注册程序
func (r *Registry) Register(p types.Program) {
	r.programs[p.Name()] = p
}

// Get 获取程序
func (r *Registry) Get(name string) (types.Program, error) {
	p, ok := r.programs[name]
	if !ok {
		return nil, fmt.Errorf("程序 '%s' 未找到", name)
	}
	return p, nil
}

// List 列出所有程序
func (r *Registry) List() []types.Program {
	programs := make([]types.Program, 0, len(r.programs))
	for _, p := range r.programs {
		programs = append(programs, p)
	}
	return programs
}

// registerAll 注册所有支持的程序
func (r *Registry) registerAll() {
	// 注册 WireGuard
	r.Register(NewWireGuard())

	// 后续可以在这里添加更多程序
	// r.Register(NewNodejs())
	// r.Register(NewPython())
}
