-- phpMyAdmin SQL Dump
-- version 4.4.13.1deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2015-12-24 20:19:01
-- 服务器版本： 5.6.27-0ubuntu1
-- PHP Version: 5.6.11-1ubuntu3.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `numa`
--

-- --------------------------------------------------------

--
-- 表的结构 `tb_admin`
--

CREATE TABLE IF NOT EXISTS `tb_admin` (
  `id` int(11) NOT NULL,
  `account` varchar(32) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `last_ip` varchar(32) NOT NULL DEFAULT '',
  `last_time` datetime NOT NULL,
  `email` varchar(100) NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_article`
--

CREATE TABLE IF NOT EXISTS `tb_article` (
  `id` int(11) NOT NULL,
  `cate_id` int(11) NOT NULL DEFAULT '0',
  `title` varchar(100) NOT NULL DEFAULT '',
  `color` varchar(7) NOT NULL DEFAULT '#00000',
  `author` varchar(15) NOT NULL DEFAULT '',
  `tags` varchar(100) NOT NULL DEFAULT '',
  `thumb_url` varchar(255) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `hits` int(11) NOT NULL DEFAULT '0',
  `post_time` datetime NOT NULL,
  `last_time` datetime NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `order` tinyint(4) NOT NULL DEFAULT '0',
  `seo_title` varchar(255) NOT NULL DEFAULT '',
  `seo_keys` varchar(255) NOT NULL DEFAULT '',
  `seo_desc` longtext NOT NULL,
  `comment_count` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_article_cate`
--

CREATE TABLE IF NOT EXISTS `tb_article_cate` (
  `id` int(11) NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT '',
  `alias` varchar(32) NOT NULL DEFAULT '',
  `tags` varchar(100) NOT NULL DEFAULT '',
  `thumb_url` varchar(255) NOT NULL DEFAULT '',
  `pid` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `order` tinyint(4) NOT NULL DEFAULT '0',
  `seo_title` varchar(255) NOT NULL DEFAULT '',
  `seo_keys` varchar(255) NOT NULL DEFAULT '',
  `seo_desc` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_auto_user`
--

CREATE TABLE IF NOT EXISTS `tb_auto_user` (
  `id` int(11) NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_baoliao`
--

CREATE TABLE IF NOT EXISTS `tb_baoliao` (
  `id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `origin` varchar(100) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `client_ip` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT '',
  `post_time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_item`
--

CREATE TABLE IF NOT EXISTS `tb_item` (
  `id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  `cate_id` int(11) NOT NULL DEFAULT '0',
  `mall_id` int(11) NOT NULL DEFAULT '0',
  `color` varchar(7) NOT NULL DEFAULT '#000000',
  `url` varchar(255) NOT NULL DEFAULT '',
  `tags` varchar(100) NOT NULL DEFAULT '',
  `remote_image_url` varchar(255) NOT NULL DEFAULT '',
  `local_image_url` varchar(255) NOT NULL DEFAULT '',
  `price` varchar(32) NOT NULL DEFAULT '',
  `old_price` varchar(32) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `user_name` varchar(32) NOT NULL DEFAULT '',
  `praise` int(11) NOT NULL DEFAULT '0',
  `bad` int(11) NOT NULL DEFAULT '0',
  `favs` int(11) NOT NULL DEFAULT '0',
  `hot` int(11) NOT NULL DEFAULT '0',
  `recommend` tinyint(4) NOT NULL DEFAULT '0',
  `hits` int(11) NOT NULL DEFAULT '0',
  `post_time` datetime NOT NULL,
  `last_time` datetime NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `order` tinyint(4) NOT NULL DEFAULT '0',
  `seo_title` varchar(255) NOT NULL DEFAULT '',
  `seo_keys` varchar(255) NOT NULL DEFAULT '',
  `seo_desc` longtext NOT NULL,
  `comment_count` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_cate`
--

CREATE TABLE IF NOT EXISTS `tb_item_cate` (
  `id` int(11) NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT '',
  `alias` varchar(32) NOT NULL DEFAULT '',
  `pid` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `order` tinyint(4) NOT NULL DEFAULT '0',
  `seo_title` varchar(255) NOT NULL DEFAULT '',
  `seo_keys` varchar(255) NOT NULL DEFAULT '',
  `seo_desc` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_comment`
--

CREATE TABLE IF NOT EXISTS `tb_item_comment` (
  `id` int(11) NOT NULL,
  `item_id` int(11) NOT NULL DEFAULT '0',
  `pid` int(11) NOT NULL DEFAULT '0',
  `client_ip` varchar(32) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `post_time` datetime NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_item_tag`
--

CREATE TABLE IF NOT EXISTS `tb_item_tag` (
  `id` int(11) NOT NULL,
  `item_id` int(11) NOT NULL DEFAULT '0',
  `tag_id` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_mall`
--

CREATE TABLE IF NOT EXISTS `tb_mall` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  `remark` varchar(100) NOT NULL DEFAULT '',
  `bimage` varchar(255) NOT NULL DEFAULT '',
  `mimage` varchar(255) NOT NULL DEFAULT '',
  `simage` varchar(255) NOT NULL DEFAULT '',
  `desc` longtext NOT NULL,
  `post_time` datetime NOT NULL,
  `last_time` datetime NOT NULL,
  `order` tinyint(4) NOT NULL DEFAULT '0',
  `seo_title` varchar(255) NOT NULL DEFAULT '',
  `seo_keys` varchar(255) NOT NULL DEFAULT '',
  `seo_desc` longtext NOT NULL,
  `comment_count` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `tb_tag`
--

CREATE TABLE IF NOT EXISTS `tb_tag` (
  `id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_admin`
--
ALTER TABLE `tb_admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `account` (`account`),
  ADD KEY `tb_admin_last_time` (`last_time`);

--
-- Indexes for table `tb_article`
--
ALTER TABLE `tb_article`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_article_cate_id` (`cate_id`),
  ADD KEY `tb_article_title` (`title`),
  ADD KEY `tb_article_author` (`author`),
  ADD KEY `tb_article_tags` (`tags`),
  ADD KEY `tb_article_post_time` (`post_time`),
  ADD KEY `tb_article_last_time` (`last_time`),
  ADD KEY `tb_article_status` (`status`);

--
-- Indexes for table `tb_article_cate`
--
ALTER TABLE `tb_article_cate`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_article_cate_tags` (`tags`);

--
-- Indexes for table `tb_auto_user`
--
ALTER TABLE `tb_auto_user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `tb_baoliao`
--
ALTER TABLE `tb_baoliao`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_baoliao_title` (`title`),
  ADD KEY `tb_baoliao_origin` (`origin`),
  ADD KEY `tb_baoliao_post_time` (`post_time`);

--
-- Indexes for table `tb_item`
--
ALTER TABLE `tb_item`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_item_cate_id` (`cate_id`),
  ADD KEY `tb_item_mall_id` (`mall_id`),
  ADD KEY `tb_item_tags` (`tags`),
  ADD KEY `tb_item_user_id` (`user_id`),
  ADD KEY `tb_item_recommend` (`recommend`),
  ADD KEY `tb_item_post_time` (`post_time`),
  ADD KEY `tb_item_last_time` (`last_time`),
  ADD KEY `tb_item_status` (`status`);

--
-- Indexes for table `tb_item_cate`
--
ALTER TABLE `tb_item_cate`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_item_cate_status` (`status`);

--
-- Indexes for table `tb_item_comment`
--
ALTER TABLE `tb_item_comment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_item_comment_item_id` (`item_id`),
  ADD KEY `tb_item_comment_post_time` (`post_time`),
  ADD KEY `tb_item_comment_status` (`status`);

--
-- Indexes for table `tb_item_tag`
--
ALTER TABLE `tb_item_tag`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_item_tag_item_id` (`item_id`),
  ADD KEY `tb_item_tag_tag_id` (`tag_id`);

--
-- Indexes for table `tb_mall`
--
ALTER TABLE `tb_mall`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_mall_post_time` (`post_time`),
  ADD KEY `tb_mall_last_time` (`last_time`);

--
-- Indexes for table `tb_tag`
--
ALTER TABLE `tb_tag`
  ADD PRIMARY KEY (`id`),
  ADD KEY `tb_tag_name` (`name`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_admin`
--
ALTER TABLE `tb_admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_article`
--
ALTER TABLE `tb_article`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_article_cate`
--
ALTER TABLE `tb_article_cate`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_auto_user`
--
ALTER TABLE `tb_auto_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_baoliao`
--
ALTER TABLE `tb_baoliao`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_item`
--
ALTER TABLE `tb_item`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_item_cate`
--
ALTER TABLE `tb_item_cate`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_item_comment`
--
ALTER TABLE `tb_item_comment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_item_tag`
--
ALTER TABLE `tb_item_tag`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_mall`
--
ALTER TABLE `tb_mall`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tb_tag`
--
ALTER TABLE `tb_tag`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
