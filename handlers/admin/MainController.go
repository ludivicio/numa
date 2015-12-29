package admin

import (
	"numa/models"
	"fmt"
)

type MainController struct {
	BaseController
}

func (m *MainController) Get() {
    fmt.Println(models.Article{})
	m.TplNames = "admin/index.html"
}
