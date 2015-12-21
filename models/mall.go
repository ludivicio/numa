package models

import (
	"github.com/astaxie/beego/orm"
  "time" 
)

/*

CREATE TABLE IF NOT EXISTS `mall` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL COMMENT '大类别的名称',
  `remark` varchar(45) DEFAULT '' COMMENT '备注',
  `bimage` varchar(45) NOT NULL COMMENT '商城logo（大图）',
  `mimage` varchar(45) DEFAULT NULL COMMENT '商城logo（中图）',
  `simage` varchar(45) DEFAULT NULL COMMENT '商城logo（小图）',
  `desc` text COMMENT '商城介绍',
  `order` tinyint(3) DEFAULT NULL COMMENT '排序',
  `seo_title` varchar(255) DEFAULT NULL COMMENT 'seo title',
  `seo_keys` varchar(255) DEFAULT NULL COMMENT 'SEO关键字',
  `seo_desc` text COMMENT 'SEO描述',
  `add_time` int(10) NOT NULL COMMENT '添加时间',
  `last_time` int(10) NOT NULL COMMENT '最后一次修改的时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品所属大分类' AUTO_INCREMENT=1 ;

*/

type Mall struct {
  Id int64
  CateId int64 `orm:"column(cate_id);index"`
  Name string `orm:"size(100)"`
  Remark string `orm:"size(100)"`
  BigImage string `orm:"column(bimage);size(255)"`
  MediumImage string `orm:"column(mimage);size(255)"`
  SmallImage string `orm:"column(simage);size(255)"`
  Desc string `orm:"type(text)"`
  PostTime time.Time `orm:"type(datetime);index"`
  LastTime time.Time `orm:"type(datetime);index"`
  Order int8 
  SeoTitle string `orm:"column(seo_title);size(255)"`
  SeoKeys string `orm:"column(seo_keys);size(255)"`
  SeoDesc string `orm:"column(seo_desc);type(text)"`
  CommentCount int32 `orm:"column(comment_count)"`
}

func (m *Mall) TableName() string {
  return TableName("mall")
}

func (m *Mall) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Mall) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mall) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mall) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Mall) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}