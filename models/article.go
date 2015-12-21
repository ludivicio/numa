package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)

/*

CREATE TABLE IF NOT EXISTS `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cate_id` tinyint(3) unsigned NOT NULL COMMENT '文章类别ID',
  `title` varchar(255) NOT NULL COMMENT '文章标题',
  `author` varchar(100) NOT NULL COMMENT '文章作者',
  `tags` varchar(45) DEFAULT NULL COMMENT '文章标签',
  `image` varchar(100) NOT NULL COMMENT '缩略图url',
  `content` text NOT NULL COMMENT '文章内容',
  `hits` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `post_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否可用 {0:不可用, 1:可用}',
  `order` tinyint(3) unsigned NOT NULL DEFAULT '255' COMMENT '排序',
  `seo_title` varchar(255) NOT NULL COMMENT 'SEO标题',
  `seo_keys` varchar(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` text NOT NULL COMMENT 'SEO描述',
  `comment_count` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表' AUTO_INCREMENT=1 ;

*/

type Article struct {
  Id int64
  CateId int64 `orm:"column(cate_id);index"`
  Title string `orm:"size(100)"`
  Color string `orm:"size(7)"`
  Author string `orm:"size(15)"`
  Tags string  `orm:"size(100)"`
  ThumbUrl string `orm:"size(100)"`
  Content string `orm:"type(text)"`
  Hits int32
  PostTime time.Time `orm:"type(datetime);index"`
  LastTime time.Time `orm:"type(datetime);index"`
  Status int8
  Order int8 
  SeoTitle string `orm:"column(seo_title);size(255)"`
  SeoKeys string `orm:"column(seo_keys);size(255)"`
  SeoDesc string `orm:"column(seo_desc);type(text)"`
  CommentCount int32 `orm:"column(comment_count)"`
}

func (m *Article) TableName() string {
  return TableName("article")
}

func (m *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}




