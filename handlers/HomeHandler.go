package handlers

import (
	"fmt"
)

type HomeHandler struct {
	BaseHandler
}

func (h *HomeHandler) Get() {
	h.Data["Website"] = "beego.me"
	h.Data["Email"] = "astaxie@gmail.com"
	h.TplNames = "home/index.tpl"
	fmt.Println("test home ")
}
