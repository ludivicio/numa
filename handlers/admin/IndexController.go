package admin

import (
	"os"
	"runtime"
)

// IndexController 后台首页
type IndexController struct {
	BaseController
}

// Index 后台首页
func (m *IndexController) Index() {
	m.Data["title"] = "首页"
	m.Data["hostname"], _ = os.Hostname()
	m.Data["goversion"] = runtime.Version()
	m.Data["os"] = runtime.GOOS
	m.Data["cpunum"] = runtime.NumCPU()
	m.Data["arch"] = runtime.GOARCH
	m.display()
}
