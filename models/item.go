package models

import (
    "github.com/astaxie/beego/orm"
    "time"  
)


/*

CREATE TABLE IF NOT EXISTS `item` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '商品标题',
  `cate_id` int(10) unsigned NOT NULL COMMENT '分类ID',
  `mall_id` int(10) unsigned NOT NULL COMMENT '商城ID',
  `tcolor` varchar(10) NOT NULL COMMENT '标题颜色',
  `url` varchar(255) NOT NULL COMMENT '购买链接',
  `price` varchar(255) DEFAULT NULL COMMENT '折扣价',
  `old_price` varchar(255) DEFAULT NULL COMMENT '原价',
  `info` text NOT NULL COMMENT '商品信息',
  `image` varchar(255) NOT NULL COMMENT '商品缩略图',
  `uid` int(10) NOT NULL COMMENT '发布者ID',
  `uname` varchar(100) NOT NULL COMMENT '发布者名称',
  `add_time` int(10) unsigned NOT NULL COMMENT '发布时间',
  `last_time` int(10) unsigned NOT NULL COMMENT '最后修改的时间',
  `good` int(10) unsigned NOT NULL COMMENT '赞',
  `bad` int(10) unsigned NOT NULL COMMENT '踩',
  `favs` int(10) unsigned NOT NULL COMMENT '收藏数',
  `hits` int(10) unsigned NOT NULL COMMENT '点击量',
  `order` tinyint(3) unsigned NOT NULL COMMENT '排序',
  `hot` tinyint(1) NOT NULL COMMENT '是否热门',
  `recommend` tinyint(1) NOT NULL COMMENT '是否推荐',
  `comments` tinyint(3) unsigned NOT NULL COMMENT '评论数',
  `localimage` tinyint(1) NOT NULL COMMENT '本地存储图片',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否可用 {0:不可用, 1:可用, 2:已过期}',
  `seo_title` varchar(255) NOT NULL COMMENT 'SEO标题',
  `seo_keys` varchar(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` text NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表' AUTO_INCREMENT=1 ;

*/

type Item struct {
  Id int64
  Title string `orm:"size(100)"`
  CateId int64 `orm:"column(cate_id);index"`
  MallId int64 `orm:"column(mall_id);index"`
  Color string `orm:"size(7)"`
  Url string `orm:"size(100)"` 
  Tags string  `orm:"size(100)"`
  RemoteThumbUrl string `orm:"column(remote_image_url);size(100)"`
  LocalThumbUrl string `orm:"column(local_image_url);size(100)"`
  Price string `orm:"size(100)"`
  OldPrice string `orm:"column(old_price);size(100)"`
  Content string `orm:"type(text)"`
  UserId int64 `orm:"column(user_id);index"`
  UserName string  `orm:"column(user_name);size(100)"`
  Praise int32 
  Bad int32
  Favs int32
  Hot int32
  Recommend int8
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