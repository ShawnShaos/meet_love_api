-- phpMyAdmin SQL Dump
-- version phpStudy 2014
-- http://www.phpmyadmin.net
--
-- 主机: localhost
-- 生成日期: 2019 �?04 �?26 �?06:38
-- 服务器版本: 5.5.53
-- PHP 版本: 5.6.27

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `test`
--

-- --------------------------------------------------------

--
-- 表的结构 `jie_lin_users`
--

CREATE TABLE IF NOT EXISTS `jie_lin_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `nickname` varchar(50) COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '头像',
  `phone` varchar(20) COLLATE utf8_bin DEFAULT NULL COMMENT '手机号',
  `email` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '邮箱',
  `create_time` datetime NOT NULL COMMENT '注册时间',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `last_login_time` datetime DEFAULT NULL COMMENT '上次登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_bin AUTO_INCREMENT=4 ;

--
-- 转存表中的数据 `jie_lin_users`
--

INSERT INTO `jie_lin_users` (`id`, `nickname`, `avatar`, `phone`, `email`, `create_time`, `login_time`, `last_login_time`) VALUES
(1, '张三', 'http://b-ssl.duitang.com/uploads/item/201706/22/20170622131955_h4eZS.thumb.700_0.jpeg', '15881630394', '92324442@qq.com', '2019-04-23 04:05:10', '2019-04-25 07:21:19', '2019-04-24 05:29:17'),
(2, '李四', 'http://img.52z.com/upload/news/image/20180823/20180823122912_66977.jpg', '1588162342', '92231242@qq.com', '2019-04-23 10:05:10', '2019-04-25 07:21:19', '2019-04-24 05:29:17'),
(3, '王五', 'http://img.52z.com/upload/news/image/20180628/20180628064705_79123.jpg', '1588124242', '922131242@qq.com', '2019-04-23 10:05:10', '2019-04-25 14:21:35', '2019-04-24 11:29:17');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
