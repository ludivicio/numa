package admin

// DataController 数据 包括爆料，标签等
type DataController struct {
	BaseController
}

func (m *DataController) Get() {
	m.display("data")
}