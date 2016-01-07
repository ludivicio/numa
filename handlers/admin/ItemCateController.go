package admin


// ItemCateController 商品分类
type ItemCateController struct {
	BaseController
}

func (m *ItemCateController) Get() {
	m.display("item-cate")
}