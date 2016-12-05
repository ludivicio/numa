package admin

import (
	"fmt"
	"numa/models"
	"strings"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ItemCateController 商品分类
type ItemCateController struct {
	BaseController
}

// ItemCateController 获取分类列表
func (m *ItemCateController) List() {

	o := orm.NewOrm()

	itemCates := []models.ItemCate{}
	dbPrefix := beego.AppConfig.String("dbprefix")

	sql := "SELECT * FROM " + dbPrefix + "item_cate"
	num, err := o.Raw(sql).QueryRows(&itemCates)
	if err == nil && num > 0 {
		for _, cate := range itemCates {
			fmt.Println(cate)
		}
	}
	m.Data["cates"] = itemCates
	m.display("item-cate")
}

//ItemCateController 处理添加分类
func (m *ItemCateController) Add() {

	if m.Ctx.Request.Method == "POST" {
		pidStr := strings.TrimSpace(m.GetString("pid"))
		name := strings.TrimSpace(m.GetString("name"))
		statusStr := strings.TrimSpace(m.GetString("status"))
		orderStr := strings.TrimSpace(m.GetString("order"))
		seoTitle := strings.TrimSpace(m.GetString("seo_title"))
		seoKeys := strings.TrimSpace(m.GetString("seo_keys"))
		seoDesc := strings.TrimSpace(m.GetString("seo_desc"))

		itemCate := models.ItemCate{Name: name, SeoTitle: seoTitle, SeoKeys: seoKeys, SeoDesc: seoDesc}

		if pid, err := strconv.Atoi(pidStr); err != nil {
			m.error(err.Error())
		} else {
			itemCate.Pid = int32(pid)
		}

		if order, err := strconv.Atoi(orderStr); err != nil {
			m.error(err.Error())
		} else {
			itemCate.Order = int8(order)
		}

		if status, err := strconv.Atoi(statusStr); err != nil {
			m.error(err.Error())
		} else {
			itemCate.Status = int8(status)
		}

		o := orm.NewOrm()
		if row, err := o.Insert(&itemCate); err != nil {
			m.error(err.Error())
		} else {
			fmt.Printf("%d", row)
		}

		fmt.Printf("%+v\n", &itemCate)
	}

	m.display("item-cate")
}
