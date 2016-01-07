package admin

// ArticleCateController 文章分类
type ArticleCateController struct {
	BaseController
}

func (m *ArticleCateController) Get() {
	m.display("article-cate")
}