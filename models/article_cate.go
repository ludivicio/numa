package models 

import (
    "github.com/astaxie/beego/orm"
)

/*

CREATE TABLE IF NOT EXISTS `article_cate` (
  `id` tinyint(3) NOT NULL,
  `name` varchar(100) NOT NULL COMMENT '类别名称',
  `alias` varchar(45) NOT NULL COMMENT '别名',
  `thumb_url` varchar(100) DEFAULT NULL COMMENT '图片URL',
  `pid` tinyint(3) unsigned NOT NULL COMMENT '父ID',
  `order` tinyint(3) unsigned NOT NULL DEFAULT '255' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否可用 {0:不可用, 1:可用}',
  `seo_title` varchar(2555) NOT NULL COMMENT 'SEO标题',
  `seo_keys` varchar(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` text NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章类别表';

*/

type ArticleCate struct {
  Id int64
  Name string `orm:"size(32)"`
  Alias string `orm:"size(32)"`
  Tags string  `orm:"size(100)"`
  ThumbUrl string `orm:"size(100)"`
  Pid int64
  Status int8
  Order int8 
  SeoTitle string `orm:"column(seo_title);size(255)"`
  SeoKeys string `orm:"column(seo_keys);size(255)"`
  SeoDesc string `orm:"column(seo_desc);type(text)"`
}

func (m *ArticleCate) TableName() string {
  return TableName("article_cate")
}

func (m *ArticleCate) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *ArticleCate) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ArticleCate) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ArticleCate) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *ArticleCate) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}