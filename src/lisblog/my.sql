-- MySQL dump 10.14  Distrib 5.5.37-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: goblog
-- ------------------------------------------------------
-- Server version	5.5.37-MariaDB-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tb_option`
--

DROP TABLE IF EXISTS `tb_option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_option` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_option`
--

LOCK TABLES `tb_option` WRITE;
/*!40000 ALTER TABLE `tb_option` DISABLE KEYS */;
INSERT INTO `tb_option` VALUES (1,'sitename','Wooya!'),(2,'siteurl','http://42.cnssuestc.org'),(3,'subtitle','42\'s blog'),(4,'pagesize','10'),(5,'keywords','Go语言,博客程序,GoBlog'),(6,'description','42的小站'),(7,'email','doomsplayer@gmail.com'),(8,'theme','default'),(9,'timezone','8'),(10,'stat','');
/*!40000 ALTER TABLE `tb_option` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_post`
--

DROP TABLE IF EXISTS `tb_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_post` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `userid` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `author` varchar(15) NOT NULL DEFAULT '' COMMENT '作者',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `color` varchar(7) NOT NULL DEFAULT '' COMMENT '标题颜色',
  `urlname` varchar(100) NOT NULL DEFAULT '' COMMENT 'url名',
  `urltype` tinyint(3) NOT NULL DEFAULT '0' COMMENT 'url访问形式',
  `excerpt` mediumtext NOT NULL COMMENT '内容摘要',
  `content` mediumtext NOT NULL COMMENT '内容',
  `tags` varchar(100) NOT NULL DEFAULT '' COMMENT '标签',
  `views` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '查看次数',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态{0:正常,1:草稿,2:回收站}',
  `posttime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  `updated` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `istop` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否置顶',
  PRIMARY KEY (`id`),
  KEY `userid` (`userid`),
  KEY `posttime` (`posttime`),
  KEY `urlname` (`urlname`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_post`
--

LOCK TABLES `tb_post` WRITE;
/*!40000 ALTER TABLE `tb_post` DISABLE KEYS */;
INSERT INTO `tb_post` VALUES (1,1,'admin','关于我','','about',1,'','<p>个人简介</p>','',0,0,'2013-12-31 10:27:49','2013-12-31 10:27:53',0),(2,1,'admin','友情链接','','links',1,'','<p><a href=\"http://www.lisijie.org\" target=\"_blank\" title=\"http://www.lisijie.org\">老李的博客</a></p>','',2,0,'2013-12-31 10:29:54','2013-12-31 10:29:54',0),(3,1,'admin','Evasi0n7 1.0.1发布 明天或再更新提高兼容','','',0,'','<p>今天凌晨 iOS 7 越狱工具 Evasi0n7 获得了更新，更新版本是 1.0.1（论坛下载）。根据 pod2g 在推特上的介绍，本次更新删除了太极助手。最近关注越狱的用户都知道，因为捆绑太极助手，越狱社区最近风波不断。</p><p><br/></p><p>许多用户对于 Evasi0n7 捆绑太极助手一事非常气愤，因为太极助手事关隐私安全问题。虽然太极助手只是针对中国用户，但是民愤仍然难平，即使 Geohot 已经证实 Evasi0n7 越狱程序中并未植入任何恶意程序。</p><p><br/></p><p>pod2g 和 Evad3rs 团队也听到了用户的“心声”，觉得必须对此采取相应措施，因此他们发布了更新。另外 pod2g 在推特上表示，Evasi0n7 将会再度迎来更新，主要是修复与 iPad 2 的兼容性问题，并表示时间“可能是明天”，即 12 月 26 日。</p><p><br/></p>',',iPhone,越狱,',1,0,'2013-12-31 10:33:41','2013-12-31 10:33:51',0),(4,1,'admin','Evasi0n发布1.0.1版越狱工具 彻底移除“太极”','','',0,'','<p>今天在twitter上宣布了最新的越狱社区消息， Evasi0n 1.0.1版本已经正式发布，之前引发轩然大波的“中国元素”- 太极助手已经彻底从工具中移除，同时他还祝大家节日快乐。</p><p><br/><img src=\"http://static.cnbetacdn.com/newsimg/2013/1225/65_1387937741.png\"/></p><p>访问网站(已经不会自动跳转中文页面):</p><p><a title=\"\" target=\"_blank\" href=\"http://evasi0n.com/\" _hover-ignore=\"1\">http://evasi0n.com/</a><br/></p><p><img src=\"http://static.cnbetacdn.com/newsimg/2013/1225/98_1387937696.png_w600.png\"/></p><p><br/></p>',',iPhone,越狱,',0,0,'2013-12-31 10:33:46','2013-12-31 10:33:49',0),(5,1,'admin','evad3rs公开信第二部分：已拒绝太极所有款项','','',0,'','<p>evad3rs于今天在越狱社区发布了公开信的第二部分，该部分对安全、盗版、以及收钱问题做了详细的解释。</p><p>亲爱的越狱社区，我们要对社会上一些额外的关注做出解释：</p><p><strong>隐私与太极</strong></p><p>首先要解释的是最重要，最令人关注的一点，是隐私。没有任何一个人数据被泄露到任何地方。当然，作为越狱社区的成员，这次的事件有损我们7年来的安全声誉。但是我们需要重申的是，除非你的设备系统语言被设置成中文，否则不会有太极软件被安装。此外，除非你单独打开了太极的应用，否则也不会有太极软件运行。</p><p>之后有谣言四起，说我们有加密的数据被发送给了装了太极的中国用户。于是我们决定作了我们最擅长的——逆向研究了太极的代码用以了解有什么数据被发送了。太极的传输数据类似于Cydia的传输数据，唯一的设备识别码是以被加密的形式传送的，类似于Cydia中使用SSL来保护用户隐私的加密形式。太极完全没有传送用户的任何隐私数据到任何地方。</p><p><strong>盗版与太极</strong></p><p>我们的所有与太极的书面协议和口头协议均以取缔。虽然太极之前向我们保证没有盗版应用，但我们并没有去仔细检查他们商店的每一个应用安装包，我们当时只是做了一个粗略的检查，并没有发现任何问题。尽管如此，后来我们接到社区的一些调查和通知后才发现了一些盗版的应用。尽管起初我们并不相信太极是故意违反我们的协议，但这次事件对我们软件开发人员以及越狱社区的声誉的影响不可小觑，我们并不会在修复了越狱工具后就忘记这次事件的，我们会铭记于心并检讨。我们已经终止了与太极的合作关系。我们对太极居然让安装了太极的越狱工具破解版运行在他们网站上感到非常失望，我们并没有给他们任何许可或者源代码。</p><p><strong>我们拒绝了所有来自太极的款项</strong></p><p>目前社会上有很多传闻说我们收取了太极的款项。我们没有收取从任何团体收取任何款项，包括太极。我们以后也将不会接受任何金钱。 我们接受的捐款都已经给了公众机构，电子前沿基金会和自由信息基础设施基金会，以帮助保护越狱作为你的合法权利。</p><p><strong>越狱的更新</strong></p><p>我们正在努力解决越狱工具的各种问题。不幸的是，现在正是节日期间，我们的团队更想多花费时间陪伴我们的家人和朋友。所以我们现在的时间很紧张也颇有压力，我们需要多一些的时间来恢复。我们会竭尽全力努力工作以解决所有的遗留问题。感谢您对我们的理解。</p><p>我们非常努力地把越狱免费提供给所有人。我们也希望您可以享受我们的越狱工具。</p><p><br/></p>',',iPhone,越狱,',1,0,'2013-12-31 10:33:50','2013-12-31 10:33:50',0);
/*!40000 ALTER TABLE `tb_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_tag`
--

DROP TABLE IF EXISTS `tb_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_tag` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '标签名',
  `count` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '使用次数',
  PRIMARY KEY (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_tag`
--

LOCK TABLES `tb_tag` WRITE;
/*!40000 ALTER TABLE `tb_tag` DISABLE KEYS */;
INSERT INTO `tb_tag` VALUES (1,'iPhone',3),(2,'越狱',3);
/*!40000 ALTER TABLE `tb_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_tag_post`
--

DROP TABLE IF EXISTS `tb_tag_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_tag_post` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tagid` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '标签id',
  `postid` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '内容id',
  `poststatus` tinyint(3) NOT NULL DEFAULT '0' COMMENT '内容状态',
  `posttime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `tagid` (`tagid`),
  KEY `postid` (`postid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_tag_post`
--

LOCK TABLES `tb_tag_post` WRITE;
/*!40000 ALTER TABLE `tb_tag_post` DISABLE KEYS */;
INSERT INTO `tb_tag_post` VALUES (1,1,22,0,'2013-12-31 10:33:46'),(2,2,22,0,'2013-12-31 10:33:46'),(3,1,21,0,'2013-12-31 10:33:49'),(4,2,21,0,'2013-12-31 10:33:49'),(5,1,20,0,'2013-12-31 10:33:51'),(6,2,20,0,'2013-12-31 10:33:51');
/*!40000 ALTER TABLE `tb_tag_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user`
--

DROP TABLE IF EXISTS `tb_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_user` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(15) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `logincount` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `lastip` varchar(15) NOT NULL DEFAULT '0' COMMENT '最后登录ip',
  `lastlogin` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `authkey` char(10) NOT NULL DEFAULT '' COMMENT '登录key',
  `active` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否激活',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user`
--

LOCK TABLES `tb_user` WRITE;
/*!40000 ALTER TABLE `tb_user` DISABLE KEYS */;
INSERT INTO `tb_user` VALUES (1,'admin','7fef6171469e80d32c0559f88b377245','admin@admin.com',7,'127.0.0.1','2014-06-08 02:15:46','',1);
/*!40000 ALTER TABLE `tb_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-06-08  2:23:17
