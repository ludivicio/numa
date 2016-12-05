package home

import (
	"fmt"
)

type DetailController struct {
	BaseController
}

func (d *DetailController) Get() {
	d.Data["Website"] = "beego.me"
	d.Data["Email"] = "astaxie@gmail.com"
	d.TplName = "home/detail.tpl"
	fmt.Println("test")
}
