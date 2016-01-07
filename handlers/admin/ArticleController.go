package admin

// ArticleController 文章
type ArticleController struct {
	BaseController
}

func (m *ArticleController) Get() {
	m.display("article")
}