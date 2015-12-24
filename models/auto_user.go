package models 

import (
    "github.com/astaxie/beego/orm"
)

/*

CREATE TABLE IF NOT EXISTS `auto_user` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统自动生成的用户';

*/

// 机器用户表
type AutoUser struct {
    Id int32
    Name string      `orm:"unique;size(32)"`
}

func (m *AutoUser) TableName() string {
  return TableName("auto_user")
}

func (m *AutoUser) Insert() error {
  if _, err := orm.NewOrm().Insert(m); err != nil {
    return err
  }
  return nil
}

func (m *AutoUser) Read(fields ...string) error {
  if err := orm.NewOrm().Read(m, fields...); err != nil {
    return err
  }
  return nil
}

func (m *AutoUser) Update(fields ...string) error {
  if _, err := orm.NewOrm().Update(m, fields...); err != nil {
    return err
  }
  return nil
}

func (m * AutoUser) Delete() error {
  if _, err := orm.NewOrm().Delete(m); err != nil {
    return err
  }
  return nil
}

func (m *AutoUser) Query() orm.QuerySeter {
  return orm.NewOrm().QueryTable(m)
}