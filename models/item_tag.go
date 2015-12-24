package models

import (
	"github.com/astaxie/beego/orm"
)

/*

CREATE TABLE IF NOT EXISTS `item_tag` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `item_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `tag_id` int(10) unsigned NOT NULL COMMENT '标签ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品标签对应表';

*/

//标签内容关系表
type ItemTag struct {
	Id         int32
	ItemId     int32 `orm:"column(item_id);index"`
	TagId     int32 `orm:"column(tag_id);index"`
}

func (m *ItemTag) TableName() string {
	return TableName("item_tag")
}

func (m *ItemTag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *ItemTag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ItemTag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ItemTag) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *ItemTag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}