package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)

/*

CREATE TABLE IF NOT EXISTS `baoliao` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NOT NULL COMMENT '标题',
  `url` VARCHAR(255) NOT NULL COMMENT '链接',
  `origin` VARCHAR(100) NOT NULL COMMENT '来源',
  `content` TEXT NOT NULL COMMENT '内容',
  `email` VARCHAR(100) NOT NULL COMMENT '爆料人邮箱',
  `post_time` datetime NOT NULL COMMENT '爆料时间',
  `client_ip` VARCHAR(32) NOT NULL COMMENT '客户端IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户爆料表' AUTO_INCREMENT=1 ;

*/

type Baoliao struct {
  Id int32
  Title string `orm:"size(100);index"`
  Url string `orm:"size(255)"`
  Origin string `orm:"size(100);index"`
  Content string `orm:"type(text)"`
  ClientIp string  `orm:"column(client_ip);size(32)"`
  Email string        `orm:"size(100)"`
  PostTime time.Time `orm:"type(datetime);index"`
}

func (m *Baoliao) TableName() string {
  return TableName("baoliao")
}

func (m *Baoliao) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Baoliao) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Baoliao) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Baoliao) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Baoliao) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}