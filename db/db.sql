-- 创建数据库
create database if not exists `db_apiserver`;

-- 创建表
use `db_apiserver`;
drop table if exists `tb_users`;
create table `tb_users` (
  `id` bigint(20) unsigned not null auto_increment,
  `username` varchar(255) not null,
  `password` varchar(255) not null,
  `createdAt` timestamp null default null,
  `updatedAt` timestamp null default null,
  `deletedAt` timestamp null default null,
  primary key (`id`),
  unique key `username` (`username`),
  key `idx_tb_users_deletedAt` (`deletedAt`)
) engine=MyISAM auto_increment=2 default charset=utf8;

-- Dumping data for table `tb_users` 插入一条记录用户名密码是:admin/admin
lock tables `tb_users` write;
INSERT INTO `tb_users` VALUES (0,'admin','$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG','2018-05-27 16:25:33','2018-05-27 16:25:33',NULL);
unlock tables;

--
-- 登录mysql:mysql -uroot -proot.然后执行：source /media/wangfeng/deepin-f/GitKraken/github/src/apiserver/db/db.sql
--

