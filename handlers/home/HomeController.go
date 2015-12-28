package home

import (
	"fmt"
)

type HomeController struct {
	BaseController
}

func (h *HomeController) Index() {
	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"
	h.TplNames = "home/index.tpl"
	fmt.Println("test home ")
}
