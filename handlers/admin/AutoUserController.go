package admin

// AutoUserController 机器用户
type AutoUserController struct {
	BaseController
}

func (m *AutoUserController) Get() {
	m.display("autouser")
}