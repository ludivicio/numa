package home

import (
	"fmt"
)

type HomeHandler struct {
	BaseHandler
}

func (h *HomeHandler) Index() {
	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"
	h.TplNames = "home/index.tpl"
	fmt.Println("test home ")
}
