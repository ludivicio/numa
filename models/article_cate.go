package models 

import (
    "github.com/astaxie/beego/orm"
)

/*

CREATE TABLE IF NOT EXISTS `article_cate` (
  `id` TINYINT(32) NOT NULL,
  `name` VARCHAR(32) NOT NULL COMMENT '类别名称',
  `alias` VARCHAR(32) NOT NULL COMMENT '别名',
  `thumb_url` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '图片URL',
  `pid` TINYINT(4) UNSIGNED NOT NULL COMMENT '父ID',
  `order` TINYINT(4) UNSIGNED NOT NULL DEFAULT 255 COMMENT '排序',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  `seo_title` VARCHAR(2555) NOT NULL COMMENT 'SEO标题',
  `seo_keys` VARCHAR(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` TEXT NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章类别表';

*/

type ArticleCate struct {
  Id int32
  Name string `orm:"size(32)"`
  Alias string `orm:"size(32)"`
  Tags string  `orm:"size(100);index"`
  ThumbUrl string `orm:"size(255)"`
  Pid int32
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