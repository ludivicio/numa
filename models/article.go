package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)

/*

CREATE TABLE IF NOT EXISTS `article` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cate_id` TINYINT(4) UNSIGNED NOT NULL COMMENT '文章类别ID',
  `title` VARCHAR(100) NOT NULL COMMENT '文章标题',
  `color` varchar(7) NOT NULL DEFAULT '#000000' COMMENT '标题颜色',
  `author` VARCHAR(15) NOT NULL COMMENT '文章作者',
  `tags` VARCHAR(100) NOT NULL COMMENT '文章标签',
  `thumb_url` VARCHAR(100) NOT NULL COMMENT '缩略图url',
  `content` TEXT NOT NULL COMMENT '文章内容',
  `hits` INT(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击数',
  `post_time` datetime NOT NULL COMMENT '添加时间',
  `last_time` datetime NOT NULL COMMENT '最后修改时间',
  `status` TINYINT(4) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  `order` TINYINT(4) UNSIGNED NOT NULL DEFAULT 255 COMMENT '排序',
  `seo_title` VARCHAR(255) NOT NULL COMMENT 'SEO标题',
  `seo_keys` VARCHAR(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` TEXT NOT NULL COMMENT 'SEO描述',
  `comment_count` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';
*/

type Article struct {
  Id int32
  CateId int32 `orm:"column(cate_id);index"`
  Title string `orm:"size(100);index"`
  Color string `orm:"size(7);default(#00000)"`
  Author string `orm:"size(15);index"`
  Tags string  `orm:"size(100);index"`
  ThumbUrl string `orm:"size(255)"`
  Content string `orm:"type(text)"`
  Hits int32
  PostTime time.Time `orm:"type(datetime);index"`
  LastTime time.Time `orm:"type(datetime);index"`
  Status int8 `orm:"index"`
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




