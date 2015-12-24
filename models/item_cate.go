package models 

import (
    "github.com/astaxie/beego/orm"
)

/*

CREATE TABLE IF NOT EXISTS `item_cate` (
  `id` TINYINT(4) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '分类名称',
  `alias` VARCHAR(32) NOT NULL COMMENT '别名',
  `pid` TINYINT(4) UNSIGNED NOT NULL COMMENT '父ID',
  `order` TINYINT(4) UNSIGNED NOT NULL COMMENT '排序',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  `seo_title` VARCHAR(255) NOT NULL COMMENT 'SEO标题',
  `seo_keys` VARCHAR(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` TEXT NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品分类表' AUTO_INCREMENT=1 ;

*/

type ItemCate struct {
  Id int32
  Name string `orm:"size(32)"`
  Alias string `orm:"size(32)"`
  Pid int32
  Status int8 `orm:"index"`
  Order int8 
  SeoTitle string `orm:"column(seo_title);size(255)"`
  SeoKeys string `orm:"column(seo_keys);size(255)"`
  SeoDesc string `orm:"column(seo_desc);type(text)"`
}

func (m *ItemCate) TableName() string {
  return TableName("item_cate")
}

func (m *ItemCate) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *ItemCate) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ItemCate) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ItemCate) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *ItemCate) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}