package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)


/*

CREATE TABLE IF NOT EXISTS `item` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NOT NULL COMMENT '商品标题',
  `cate_id` INT(4) UNSIGNED NOT NULL COMMENT '分类ID',
  `mall_id` INT(4) UNSIGNED NOT NULL COMMENT '商城ID',
  `color` VARCHAR(7) NOT NULL DEFAULT '#000000' COMMENT '标题颜色',
  `url` VARCHAR(255) NOT NULL COMMENT '购买链接',
  `price` VARCHAR(10) NOT NULL COMMENT '折扣价',
  `old_price` VARCHAR(10) NOT NULL COMMENT '原价',
  `content` TEXT NOT NULL COMMENT '商品信息',
  `remote_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '商品缩略图',
  `local_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '本地商品缩略图',
  `user_id` INT(10) NOT NULL COMMENT '发布者ID',
  `user_name` VARCHAR(32) NOT NULL COMMENT '发布者名称',
  `post_time` datetime NOT NULL COMMENT '发布时间',
  `last_time` datetime NOT NULL COMMENT '最后修改的时间',
  `praise` INT(10) UNSIGNED NOT NULL COMMENT '赞',
  `bad` INT(10) UNSIGNED NOT NULL COMMENT '踩',
  `favs` INT(10) UNSIGNED NOT NULL COMMENT '收藏数',
  `hits` INT(10) UNSIGNED NOT NULL COMMENT '点击量',
  `order` TINYINT(4) UNSIGNED NOT NULL COMMENT '排序',
  `hot` TINYINT(4) NOT NULL COMMENT '是否热门',
  `recommend` TINYINT(4) NOT NULL COMMENT '是否推荐',
  `comment_count` TINYINT(10) UNSIGNED NOT NULL COMMENT '评论数',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用, 2:已过期}',
  `seo_title` VARCHAR(255) NOT NULL COMMENT 'SEO标题',
  `seo_keys` VARCHAR(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` TEXT NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表' AUTO_INCREMENT=1 ;

*/

type Item struct {
  Id int32
  Title string `orm:"size(100)"`
  CateId int32 `orm:"column(cate_id);index"`
  MallId int32 `orm:"column(mall_id);index"`
  Color string `orm:"size(7);default(#000000)"`
  Url string `orm:"size(255)"` 
  Tags string  `orm:"size(100);index"`
  RemoteThumbUrl string `orm:"column(remote_image_url);size(255)"`
  LocalThumbUrl string `orm:"column(local_image_url);size(255)"`
  Price string `orm:"size(32)"`
  OldPrice string `orm:"column(old_price);size(32)"`
  Content string `orm:"type(text)"`
  UserId int32 `orm:"column(user_id);index"`
  UserName string  `orm:"column(user_name);size(32)"`
  Praise int32 
  Bad int32
  Favs int32
  Hot int32
  Recommend int8 `orm:"index"`
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

func (m *Item) TableName() string {
  return TableName("item")
}

func (m *Item) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Item) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Item) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Item) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Item) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}