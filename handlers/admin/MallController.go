package admin

// MallController 商城
type MallController struct {
	BaseController
}

func (m *MallController) Get() {
	m.display("mall")
}