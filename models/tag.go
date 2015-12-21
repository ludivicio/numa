package models 

import (
	"github.com/astaxie/beego/orm"
)

/*

CREATE TABLE IF NOT EXISTS `tag` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '标签名',
  `count` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '使用次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签表' AUTO_INCREMENT=1 ;

*/

//标签表
type Tag struct {
	Id    int64
	Name  string `orm:"size(20);index"`
	Count int64
}

func (m *Tag) TableName() string {
	return TableName("tag")
}

func (m *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}