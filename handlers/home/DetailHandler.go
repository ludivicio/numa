package home

import (
	"fmt"
)

type DetailHandler struct {
	BaseHandler
}

func (d *DetailHandler) Get() {
	d.Data["Website"] = "beego.me"
	d.Data["Email"] = "astaxie@gmail.com"
	d.TplNames = "home/detail.tpl"
	fmt.Println("test")
}
