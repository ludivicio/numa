package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)

/*

CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(16) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '管理员密码',
  `last_ip` varchar(15) NOT NULL COMMENT '最后登录的ip地址',
  `last_time` int(10) NOT NULL DEFAULT '0' COMMENT '最后登录的时间',
  `email` varchar(45) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '账号是否可用 {0:不可用, 1:可用}',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员信息表' AUTO_INCREMENT=1 ;


*/

// 管理员表
type Admin struct {
    Id int64
    Account string      `orm:"unique;size(32)"`
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

