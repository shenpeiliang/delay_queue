-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.6.17 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win32
-- HeidiSQL 版本:                  8.0.0.4396
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 gin 的数据库结构
DROP DATABASE IF EXISTS `gin`;
CREATE DATABASE IF NOT EXISTS `gin` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `gin`;


-- 导出  表 gin.hs_task 结构
DROP TABLE IF EXISTS `hs_task`;
CREATE TABLE IF NOT EXISTS `hs_task` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID编号',
  `queue_index` varchar(120) NOT NULL DEFAULT '' COMMENT '执行任务的队列索引',
  `queue_slot` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '执行任务的队列槽位',
  `cycle_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '执行任务的循环次数',
  `retry_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '执行任务的重试次数',
  `dateline` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '任务操作时间',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '任务创建时间',
  `notify_url` varchar(255) NOT NULL DEFAULT '' COMMENT '请求地址',
  `request_method` varchar(4) NOT NULL DEFAULT '' COMMENT '请求方法',
  `notify_param` varchar(255) NOT NULL DEFAULT '' COMMENT '请求参数',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '任务状态：0待处理；1处理 - 成功；2处理失败-待处理；3处理失败',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='队列任务';

-- 正在导出表  gin.hs_task 的数据：~8 rows (大约)
/*!40000 ALTER TABLE `hs_task` DISABLE KEYS */;
REPLACE INTO `hs_task` (`id`, `queue_index`, `queue_slot`, `cycle_num`, `retry_num`, `dateline`, `add_time`, `notify_url`, `request_method`, `notify_param`, `state`) VALUES
	(1, 'c05d5ec4-b290-4b43-b985-10928955fdba', 10, 2, 0, 1592531958, 1592531602, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592531732', 1),
	(2, '1ca0d561-3605-4617-b021-5e9b332ceeee', 10, 0, 0, 1592531673, 1592531602, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592531612', 1),
	(3, '21336bec-c010-41f7-a33b-ad3b6a2d0870', 10, 2, 0, 1592539039, 1592538908, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592539038', 1),
	(4, '1c7a9e31-9502-4594-816f-32442053db56', 10, 0, 0, 1592538980, 1592538908, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592538918', 1),
	(5, 'd2fcdc1a-e8c4-4c88-b5ca-14855c44e303', 10, 0, 0, 1592616569, 1592616557, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592616567', 1),
	(6, '036555f8-8702-4559-99a8-ddc840b38af3', 10, 2, 0, 1592617397, 1592616557, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592616687', 1),
	(7, '56a16f90-f761-48fe-a264-b0bef15f72dc', 10, 2, 0, 1592617397, 1592617206, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592617336', 1),
	(8, 'cb2523db-ca8b-4b53-821c-4a9c64be0ae4', 10, 0, 0, 1592617217, 1592617206, 'http://sksystem.sk0.com/queue/vip', 'GET', 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1592617216', 1),
	(9, 'd0fe1f30-182e-4f89-b0fd-54e0e025b57a', 10, 0, 0, 1592880491, 1592880365, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T10:45:54.087855+08:00","method_name":"","notify_param":""}', 1),
	(10, 'b21cd0c1-99e6-4c76-b5b6-6b4a6b38fdd2', 17, 1, 0, 1592880554, 1592880409, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T10:47:54.1728598+08:00","method_name":"","notify_param":""}', 1),
	(11, 'e57a1988-0f3d-41fc-9860-26136d9900b9', 8, 0, 0, 1592880725, 1592880477, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T10:48:05.9523972+08:00","method_name":"","notify_param":""}', 1),
	(12, '749ddc53-538d-4c7d-a12d-9979c04d0c2f', 9, 2, 0, 1592880545, 1592880477, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T10:50:06.0464026+08:00","method_name":"","notify_param":""}', 1),
	(13, '448fc543-e572-4130-92bd-3860eeb5135f', 56, 8, 0, 1592883351, 1592882994, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T11:38:50+08:00","method_name":"","notify_param":""}', 1),
	(14, '375ea792-b8a9-4ebe-b688-016e314fe377', 9, 2, 0, 1592883064, 1592882994, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T11:32:03.8484126+08:00","method_name":"","notify_param":""}', 1),
	(15, '72b8c800-bf83-46df-bfda-00164ab64bb5', 9, 0, 0, 1592883005, 1592882994, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T11:30:03.7754084+08:00","method_name":"","notify_param":""}', 1),
	(16, '12bbfac9-5eef-45c5-bc08-21ae387d0585', 9, 0, 0, 1592891777, 1592891766, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T13:56:15.5117637+08:00","method_name":"","notify_param":""}', 1),
	(17, 'aae4c32a-76df-4202-a368-3b675ae51617', 9, 2, 0, 1592891896, 1592891766, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T13:58:15.6167697+08:00","method_name":"","notify_param":""}', 1),
	(18, '279952c5-25d5-442f-90c6-cd0982e871fe', 50, 4, 0, 1592892433, 1592891830, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:02:00+08:00","method_name":"","notify_param":""}', 1),
	(19, '4d779044-9583-4ae5-9b59-8bc7f6a37973', 48, 7, 0, 1592892731, 1592891832, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:05:00+08:00","method_name":"","notify_param":""}', 1),
	(20, 'fe1e59e6-3773-4069-85b7-b5b78e10a0f2', 44, 8, 0, 1592892607, 1592891836, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:06:00+08:00","method_name":"","notify_param":""}', 1),
	(21, '9e1a1355-0de5-4015-b6d7-c45436c99053', 10, 0, 0, 1592893053, 1592893018, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:17:08.0854068+08:00","method_name":"","notify_param":""}', 1),
	(22, 'f1e6fb32-e3a1-4aba-8bf7-19aab6b0331a', 9, 2, 0, 1592895736, 1592895430, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:59:19.4943315+08:00","method_name":"","notify_param":""}', 1),
	(23, '4081e274-59d1-4949-a2a5-652408f3faa2', 9, 0, 0, 1592895441, 1592895430, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:57:19.4113268+08:00","method_name":"","notify_param":""}', 1),
	(24, 'c1f607da-bb22-4b94-b55b-f6ca8d82c6f7', 10, 0, 0, 1592895679, 1592895474, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:58:04.2618921+08:00","method_name":"","notify_param":""}', 1),
	(25, '33de6405-4e5f-4d6b-ab1a-566fce67da3b', 9, 0, 0, 1592895500, 1592895476, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:58:05.8819847+08:00","method_name":"","notify_param":""}', 1),
	(26, '97f66d2f-d123-4ac8-a5bf-344384b2f7b3', 9, 0, 0, 1592895500, 1592895477, 'http://sksystem.sk0.com/queue/vip', 'GET', '{"notify_url":"http://sksystem.sk0.com/queue/vip","plan_time":"2020-06-23T14:58:06.5760244+08:00","method_name":"","notify_param":""}', 1);
/*!40000 ALTER TABLE `hs_task` ENABLE KEYS */;


-- 导出  表 gin.hs_task_log 结构
DROP TABLE IF EXISTS `hs_task_log`;
CREATE TABLE IF NOT EXISTS `hs_task_log` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID编号',
  `queue_index` varchar(120) NOT NULL DEFAULT '' COMMENT '执行任务的队列索引',
  `response` text NOT NULL COMMENT '响应结果',
  `curl` text NOT NULL COMMENT 'curl可重发命令',
  `dateline` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '任务操作时间',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '任务状态：0待处理；1处理 - 成功；2处理失败-待处理；3处理失败',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='队列任务记录';

-- 正在导出表  gin.hs_task_log 的数据：~8 rows (大约)
/*!40000 ALTER TABLE `hs_task_log` DISABLE KEYS */;
REPLACE INTO `hs_task_log` (`id`, `queue_index`, `response`, `curl`, `dateline`, `state`) VALUES
	(1, '1ca0d561-3605-4617-b021-5e9b332ceeee', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592531673, 1),
	(2, 'c05d5ec4-b290-4b43-b985-10928955fdba', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592531958, 1),
	(3, '1c7a9e31-9502-4594-816f-32442053db56', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592538980, 1),
	(4, '21336bec-c010-41f7-a33b-ad3b6a2d0870', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592539039, 1),
	(5, 'd2fcdc1a-e8c4-4c88-b5ca-14855c44e303', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592616569, 1),
	(6, 'cb2523db-ca8b-4b53-821c-4a9c64be0ae4', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592617217, 1),
	(7, '036555f8-8702-4559-99a8-ddc840b38af3', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592617397, 1),
	(8, '56a16f90-f761-48fe-a264-b0bef15f72dc', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592617397, 1),
	(9, 'd0fe1f30-182e-4f89-b0fd-54e0e025b57a', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592880491, 1),
	(10, '749ddc53-538d-4c7d-a12d-9979c04d0c2f', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592880545, 1),
	(11, 'b21cd0c1-99e6-4c76-b5b6-6b4a6b38fdd2', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592880554, 1),
	(12, 'e57a1988-0f3d-41fc-9860-26136d9900b9', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592880725, 1),
	(13, '72b8c800-bf83-46df-bfda-00164ab64bb5', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592883005, 1),
	(14, '375ea792-b8a9-4ebe-b688-016e314fe377', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592883064, 1),
	(15, '448fc543-e572-4130-92bd-3860eeb5135f', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592883351, 1),
	(16, '12bbfac9-5eef-45c5-bc08-21ae387d0585', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592891777, 1),
	(17, 'aae4c32a-76df-4202-a368-3b675ae51617', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592891896, 1),
	(18, '279952c5-25d5-442f-90c6-cd0982e871fe', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592892433, 1),
	(19, 'fe1e59e6-3773-4069-85b7-b5b78e10a0f2', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592892607, 1),
	(20, '4d779044-9583-4ae5-9b59-8bc7f6a37973', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592892731, 1),
	(21, '9e1a1355-0de5-4015-b6d7-c45436c99053', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592893053, 1),
	(22, '4081e274-59d1-4949-a2a5-652408f3faa2', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592895440, 1),
	(23, '97f66d2f-d123-4ac8-a5bf-344384b2f7b3', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592895500, 1),
	(24, '33de6405-4e5f-4d6b-ab1a-566fce67da3b', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592895500, 1),
	(25, 'c1f607da-bb22-4b94-b55b-f6ca8d82c6f7', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592895679, 1),
	(26, 'f1e6fb32-e3a1-4aba-8bf7-19aab6b0331a', '状态码： 200 响应信息： {"success":true}', 'curl -G   -d \'\' http://sksystem.sk0.com/queue/vip', 1592895736, 1);
/*!40000 ALTER TABLE `hs_task_log` ENABLE KEYS */;


-- 导出  表 gin.hs_user 结构
DROP TABLE IF EXISTS `hs_user`;
CREATE TABLE IF NOT EXISTS `hs_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID编号',
  `is_super` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是超级管理员',
  `uname` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `real_name` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `tel` varchar(20) NOT NULL DEFAULT '' COMMENT '电话',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '注册时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_user_name` (`uname`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='会员信息';

-- 正在导出表  gin.hs_user 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `hs_user` DISABLE KEYS */;
REPLACE INTO `hs_user` (`id`, `is_super`, `uname`, `email`, `password`, `real_name`, `tel`, `add_time`) VALUES
	(1, 0, 'demo', '674894243@qq.com', '$2a$10$0GoELb6Yjq.PYNrvxUC5e.YYVu6NS87GfWPV9Gx2atf6XQY9ENzse', '哈利', '13557113347', 1589965856),
	(10, 1, '管理员', '674894243@qq.com', '$2a$10$4ah5vsG7qGRyVVZtuC9D5OMSilylxT3Y7Ne1g48KoG5eY5rgcBRPC', '沈', '13557113347', 1590462370),
	(12, 0, '管理员007', '674894243@qq.com', '$2a$10$DiApQtno1p8Lz8baaMqJ.uE4l44itHEltGeAy2FysyeHGMQKmDF.i', '水水', '12345678910', 1591941857),
	(13, 0, '试客0202', '2172592393@qq.com', '$2a$10$TRF6W7HGq7r1UoOZ7L4VR.n6CzbSLWO4Egl9mulhlP1zMm4Xi9pSu', '我的号', '13557113347', 1592374505);
/*!40000 ALTER TABLE `hs_user` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
