-- MySQL dump 10.13  Distrib 5.7.26, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: go_study
-- ------------------------------------------------------
-- Server version	5.7.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` varchar(50) NOT NULL DEFAULT '' COMMENT '全局唯一标识',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT 'name',
  `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_uuid_uindex` (`uuid`),
  UNIQUE KEY `user_name_uindex` (`name`),
  KEY `user_created_at_index` (`created_at`),
  KEY `user_updated_at_index` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'d8aae903-d52f-4d1c-b84a-ac5ea18d0f8d','aaaa','','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(2,'76c2e0c4-3bed-4d6e-be24-3fc1c289e6d1','哈哈哈','','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(3,'ed27d954-2f1c-4243-a424-1d9f483f251a','哦哦哦','','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(4,'c08f0481-8ee6-4420-8a4c-dc8971fd07ff','卡扣','','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(5,'e5fa007b-c838-44b6-820d-2f8ab5629e4e','卡扣2','','2022-04-27 15:14:00','2022-04-27 15:14:00','2022-04-22 13:45:52'),(6,'a8044564-86e5-4f33-8bea-bd4014ee7ae1','略略','','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(7,'12a33124-bba7-4f1b-b946-fccf642721cb','去玩儿','$2a$04$ecYg4N/M8IXiXlGsJ6/YQO81MK3XjO/A6dG5Mw78pBTpouPTQjZsi','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(8,'1ecf1d25-5f0e-43bb-8f7b-d8af95fe9bf5','去玩儿22','$2a$04$V07.KkqXPfqWeCkvurGDte1MaiXo9LWmaijv.cZEd6flY/3akuU9u','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(9,'dfe19073-956b-4a18-acae-d340a547dcd2','123456','$2a$04$J1TxVXKixwUOoKIr3FoY3uYG9csggHd/PSITxxpz3GnRMVnCCuwFm','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL),(10,'cb12a069-ec9e-45d3-a6f7-c605fe6561ad','111111','$2a$04$.svRq5tPttaWlXKGQEFiBuLJSA4bcsKMDw/LVPfbwSid/nZ5LZyG.','2022-04-27 15:14:00','2022-04-27 15:14:00',NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-27 23:20:17
