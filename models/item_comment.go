package models 

import (
    "github.com/astaxie/beego/orm"
    "time"  
)

/*

CREATE TABLE IF NOT EXISTS `item_comment` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `item_id` INT(10) UNSIGNED NOT NULL COMMENT '商品ID',
  `pid` INT(10) UNSIGNED NOT NULL COMMENT '父ID',
  `client_ip` VARCHAR(32) NOT NULL COMMENT '发布者的IP',
  `content` TEXT NOT NULL COMMENT '评论内容',
  `post_time` datetime NOT NULL COMMENT '评论时间',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品评论表' AUTO_INCREMENT=1 ;

*/

type ItemComment struct {
  Id int32
  ItemId int32 `orm:"column(item_id);index"`
  Pid int32
  ClientIp string  `orm:"column(client_ip);size(32)"`
  Content string `orm:"type(text)"`
  PostTime time.Time `orm:"type(datetime);index"`
  Status int8 `orm:"index"`
}

func (m *ItemComment) TableName() string {
  return TableName("item_comment")
}

func (m *ItemComment) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *ItemComment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ItemComment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *ItemComment) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *ItemComment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
