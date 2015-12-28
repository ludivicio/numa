package admin

import (
	"numa/models"
	"fmt"
)

type MainController struct {
	BaseController
}

func (m *MainController) Get() {
	m.Data["Website"] = "beego.me"
	m.Data["Email"] = "astaxie@gmail.com"
	m.TplNames = "index.tpl"
	fmt.Println("test main")
	fmt.Println(new(models.Article))
}
