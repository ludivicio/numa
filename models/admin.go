package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)

/*

CREATE TABLE IF NOT EXISTS `admin` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `account` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '账号',
  `password` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '管理员密码',
  `last_ip` VARCHAR(32) NOT NULL COMMENT '最后登录的ip地址',
  `last_time` datetime NOT NULL COMMENT '最后登录的时间',
  `email` VARCHAR(100) NULL COMMENT '邮箱',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '账号是否可用 {0:不可用, 1:可用}',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `account` (`account` ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员信息表';


*/

// 管理员表
type Admin struct {
    Id int32
    Account string      `orm:"unique;size(32);index;"`
    Password string     `orm:"size(32)"`
    LastIp string       `orm:"size(32)"`
    LastTime time.Time  `orm:"type(datatime);index"`
    Email string        `orm:"size(100)"`
    Status int8
}


func (m *Admin) TableName() string {
  return TableName("admin")
}

func (m *Admin) Insert() error {
  if _, err := orm.NewOrm().Insert(m); err != nil {
    return err
  }
  return nil
}

func (m *Admin) Read(fields ...string) error {
  if err := orm.NewOrm().Read(m, fields...); err != nil {
    return err
  }
  return nil
}

func (m *Admin) Update(fields ...string) error {
  if _, err := orm.NewOrm().Update(m, fields...); err != nil {
    return err
  }
  return nil
}

func (m * Admin) Delete() error {
  if _, err := orm.NewOrm().Delete(m); err != nil {
    return err
  }
  return nil
}

func (m *Admin) Query() orm.QuerySeter {
  return orm.NewOrm().QueryTable(m)
}

