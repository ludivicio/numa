package admin


// ItemController 商品
type ItemController struct {
	BaseController
}

func (m *ItemController) Get() {
	m.display("item")
}


