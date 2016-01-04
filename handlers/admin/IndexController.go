package admin

import (
	"numa/models"
	"fmt"
)

type IndexController struct {
	BaseController
}

func (m *IndexController) Get() {
    fmt.Println(models.Article{})
	m.TplNames = "admin/index.html"
}
