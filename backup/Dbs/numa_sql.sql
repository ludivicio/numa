SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

CREATE SCHEMA IF NOT EXISTS `numa` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ;
USE `numa` ;

-- -----------------------------------------------------
-- Table `numa`.`ka_admin`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_admin` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `account` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '账号',
  `password` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '管理员密码',
  `last_ip` VARCHAR(32) NOT NULL COMMENT '最后登录的ip地址',
  `last_time` datetime NOT NULL COMMENT '最后登录的时间',
  `email` VARCHAR(100) NULL COMMENT '邮箱',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '账号是否可用 {0:不可用, 1:可用}',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `account` (`account` ASC))
ENGINE = MyISAM
COMMENT = '管理员信息表';

--
-- 转存表中的数据 `tb_admin`
--
INSERT INTO `ka_admin` (`id`, `account`, `password`, `last_ip`, `role_id`, `last_time`, `email`, `status`) VALUES
(1, 'admin', '21232f297a57a5a743894a0e4a801fc3', '127.0.0.1', 1399293326, 'lurma@qq.com', 1);
-- --------------------------------------------------------

-- -----------------------------------------------------
-- Table `numa`.`tb_article`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_article` (
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
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '文章表';


-- -----------------------------------------------------
-- Table `numa`.`tb_article_cate`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_article_cate` (
  `id` TINYINT(4) NOT NULL,
  `name` VARCHAR(32) NOT NULL COMMENT '类别名称',
  `alias` VARCHAR(32) NOT NULL COMMENT '别名',
  `thumb_url` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '图片URL',
  `pid` TINYINT(4) UNSIGNED NOT NULL COMMENT '父ID',
  `order` TINYINT(4) UNSIGNED NOT NULL DEFAULT 255 COMMENT '排序',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  `seo_title` VARCHAR(2555) NOT NULL COMMENT 'SEO标题',
  `seo_keys` VARCHAR(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` TEXT NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '文章类别表';


-- -----------------------------------------------------
-- Table `numa`.`tb_auto_user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_auto_user` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '名称',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '系统自动生成的用户';


-- -----------------------------------------------------
-- Table `numa`.`tb_baoliao`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_baoliao` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) NOT NULL COMMENT '标题',
  `url` VARCHAR(255) NOT NULL COMMENT '链接',
  `origin` VARCHAR(100) NOT NULL COMMENT '来源',
  `content` TEXT NOT NULL COMMENT '内容',
  `email` VARCHAR(100) NOT NULL COMMENT '爆料人邮箱',
  `post_time` datetime NOT NULL COMMENT '爆料时间',
  `client_ip` VARCHAR(32) NOT NULL COMMENT '客户端IP',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '用户爆料表';

-- -----------------------------------------------------
-- Table `numa`.`tb_item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_item` (
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
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '商品表';


-- -----------------------------------------------------
-- Table `numa`.`tb_item_cate`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_item_cate` (
  `id` TINYINT(4) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL COMMENT '分类名称',
  `alias` VARCHAR(32) NOT NULL COMMENT '别名',
  `pid` TINYINT(4) UNSIGNED NOT NULL COMMENT '父ID',
  `order` TINYINT(4) UNSIGNED NOT NULL COMMENT '排序',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  `seo_title` VARCHAR(255) NOT NULL COMMENT 'SEO标题',
  `seo_keys` VARCHAR(255) NOT NULL COMMENT 'SEO关键字',
  `seo_desc` TEXT NOT NULL COMMENT 'SEO描述',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '商品分类表';


-- -----------------------------------------------------
-- Table `numa`.`tb_item_comment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`tb_item_comment` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `item_id` INT(10) UNSIGNED NOT NULL COMMENT '商品ID',
  `pid` INT(10) UNSIGNED NOT NULL COMMENT '父ID',
  `client_ip` VARCHAR(32) NOT NULL COMMENT '发布者的IP',
  `content` TEXT NOT NULL COMMENT '评论内容',
  `post_time` datetime NOT NULL COMMENT '评论时间',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否可用 {0:不可用, 1:可用}',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '商品评论表';


-- -----------------------------------------------------
-- Table `numa`.`ka_tag`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`ka_tag` (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL COMMENT '标签名',
  PRIMARY KEY (`id`))
ENGINE = MyISAM
COMMENT = '标签表';


-- -----------------------------------------------------
-- Table `numa`.`ka_item_tag`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`ka_item_tag` (
  `item_id` INT(10) UNSIGNED NOT NULL COMMENT '商品ID',
  `tag_id` INT(10) UNSIGNED NOT NULL COMMENT '标签ID',
  UNIQUE INDEX `item_id_UNIQUE` (`item_id` ASC),
  UNIQUE INDEX `tag_id_UNIQUE` (`tag_id` ASC))
ENGINE = MyISAM
COMMENT = '商品标签对应表';





-- -----------------------------------------------------
-- Table `numa`.`ka_item_from`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `numa`.`ka_item_from` (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL COMMENT '大类别的名称',
  `remark` VARCHAR(45) NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC))
ENGINE = MyISAM
COMMENT = '商品所属大分类';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
