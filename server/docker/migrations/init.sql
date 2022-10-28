-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: localhost    Database: kartenspielen
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `_prisma_migrations`
--

DROP TABLE IF EXISTS `_prisma_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `_prisma_migrations` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `checksum` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `finished_at` datetime(3) DEFAULT NULL,
  `migration_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `logs` text COLLATE utf8mb4_unicode_ci,
  `rolled_back_at` datetime(3) DEFAULT NULL,
  `started_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `applied_steps_count` int unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `_prisma_migrations`
--

LOCK TABLES `_prisma_migrations` WRITE;
/*!40000 ALTER TABLE `_prisma_migrations` DISABLE KEYS */;
INSERT INTO `_prisma_migrations` VALUES ('15386133-900e-4691-a7ca-4868040f4c1e','632c2d62ff1dd52f84b8a7b1411ec1a1d44824acc22ed7d1ed0cce8f5390635a','2022-10-18 15:08:04.642','20221015062001_',NULL,NULL,'2022-10-18 15:08:04.404',1),('1b52d426-996a-489c-adc9-c3c8b17e7fe9','3b16f8a221fbd890612885795ea4122eef6b6a83bc5a0582a6f2e18a3ce8dffe','2022-10-18 15:08:04.809','20221015144527_',NULL,NULL,'2022-10-18 15:08:04.743',1),('442f98d8-fc0b-4ffa-a2f2-b17e5b06677b','7214217aa0fd6ae30acf2535ca631e7df3f65f5c028a8796cb99aba3ae6798d8','2022-10-18 15:08:04.826','20221018143857_',NULL,NULL,'2022-10-18 15:08:04.811',1),('47465636-9d91-4af2-9dde-fbe09425c981','32945487707fa17e26beb73ef11be707ffa56e721cab28323925d4ff781ee012','2022-10-18 15:08:04.741','20221015121648_',NULL,NULL,'2022-10-18 15:08:04.676',1),('5f4434b4-fe12-44f1-b04c-b15addeca449','ca6678d012115339d158b1950f32c94b25143b52fb8e8ecd31b5b2b03e63ffbb','2022-10-18 15:08:06.262','20221018150806_',NULL,NULL,'2022-10-18 15:08:06.251',1),('7accd779-118e-41a8-b203-d81e463153b0','2b2d6635a6c53053e6b5427b27f344674e020fdca53c2a35eb7585c27377bc17','2022-10-18 15:08:04.675','20221015103318_',NULL,NULL,'2022-10-18 15:08:04.643',1),('af57a7fc-8c86-486e-84b3-548f664f5b19','12009b48b25fba8eabe839606e57ebec72612877f9a260600836578aeb049338','2022-10-18 15:08:04.118','20221010075210_init',NULL,NULL,'2022-10-18 15:08:04.104',1),('ca17e0aa-b45c-4731-a9e3-e5b9dcdbb99b','5c1c2ce6043d650f0f877241385b17955e969d79256c3e57c5ed5a1691de75cb','2022-10-18 15:08:04.332','20221015042324_',NULL,NULL,'2022-10-18 15:08:04.301',1),('fc71f95b-65fa-484f-8ec6-c457dc7fefcf','390c4907fb561dd4dcb2a2d318ee3465a8f2912eaca8b966f71ec18b1331fced','2022-10-18 15:08:04.299','20221015041437_',NULL,NULL,'2022-10-18 15:08:04.119',1),('fcb76cf2-875b-4009-a08b-1600b05b81a5','b273e4b56f50f7771656d972e194131523e8c4e3e1bda70c368b651ff5ff85d6','2022-10-18 15:08:04.402','20221015042610_',NULL,NULL,'2022-10-18 15:08:04.334',1);
/*!40000 ALTER TABLE `_prisma_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `card`
--

DROP TABLE IF EXISTS `card`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `card` (
  `id` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tableId` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `active` tinyint(1) NOT NULL,
  `createdAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updatedAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `playerId` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `Cards_tableId_fkey` (`tableId`),
  KEY `Cards_playerId_fkey` (`playerId`),
  CONSTRAINT `Cards_playerId_fkey` FOREIGN KEY (`playerId`) REFERENCES `player` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `Cards_tableId_fkey` FOREIGN KEY (`tableId`) REFERENCES `table` (`key`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `card`
--

LOCK TABLES `card` WRITE;
/*!40000 ALTER TABLE `card` DISABLE KEYS */;
/*!40000 ALTER TABLE `card` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `game`
--

DROP TABLE IF EXISTS `game`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `createdAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updatedAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `extraFields` json NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Games_name_key` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game`
--

LOCK TABLES `game` WRITE;
/*!40000 ALTER TABLE `game` DISABLE KEYS */;
INSERT INTO `game` VALUES (1,'BLACK JACK','21を目指して勝利しろ!!','blackjack.png','2022-10-15 20:00:00.000','2022-10-15 20:00:00.000','{\"turn\": \"\"}');
/*!40000 ALTER TABLE `game` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `player` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `passwordHash` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `userAuthToken` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `createdAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updatedAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `roleId` int NOT NULL,
  `tableId` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Players_username_key` (`username`),
  UNIQUE KEY `Players_userAuthToken_key` (`userAuthToken`),
  KEY `Players_roleId_fkey` (`roleId`),
  KEY `Players_tableId_fkey` (`tableId`),
  CONSTRAINT `Players_roleId_fkey` FOREIGN KEY (`roleId`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `Players_tableId_fkey` FOREIGN KEY (`tableId`) REFERENCES `table` (`key`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player`
--

LOCK TABLES `player` WRITE;
/*!40000 ALTER TABLE `player` DISABLE KEYS */;
INSERT INTO `player` VALUES (1,'lnoueryo','$2b$10$wm1sbZQWJ7gP2pcpww.6AOJVX4wLlzScLGxx3rUWAvkjFPiN9LvrS','8f7d1327-60a4-4dab-8718-5ddc57a0b20d','2022-10-18 15:11:05.076','2022-10-23 10:46:27.444',2,NULL),
(2,'player1','$2b$10$cB3mxq5J0TBeeWQbUY19yuwrrUVJZpV30TbAj7j1GKOT9Ap5Zn7k6','8c41b552-1c87-4bfd-8968-0f20f1d3fbb4','2022-10-19 10:42:02.915','2022-10-22 04:28:21.696',2,NULL),
(3,'player2','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','62290da-502f-4344-8a0d-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(4,'player3','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0da-50f-4344-8a0d-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(5,'player4','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0da-502f-4344-8ad-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(6,'player5','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0da-502f-4344-8a0d-aaf06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(7,'player6','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0da-502f-4344-8a0d-afd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(8,'player7','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0da-502f-4344-a0d-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(9,'player8','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0da-502f-344-8a0d-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(10,'player9','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','6229e0a-502f-4344-8a0d-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,NULL),
(11,'player10','$2b$10$invhOj9HHRLmbe2DRuIQyeTYCGqgYzmUW4QO.THoRdQr2jCjToXem','629e0da-502f-4344-8a0d-aafd06724406','2022-10-19 13:18:16.276','2022-10-20 14:29:53.273',2,'782087e8-1ea6-408d-b792-75ce1b60c155');
/*!40000 ALTER TABLE `player` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Roles_name_key` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'ADMIN'),(2,'USER');
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `table`
--

DROP TABLE IF EXISTS `table`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `table` (
  `key` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `gameId` int NOT NULL,
  `adminId` int NOT NULL,
  `limit` int NOT NULL,
  `createdAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updatedAt` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `start` tinyint(1) NOT NULL,
  `extraFields` json NOT NULL,
  PRIMARY KEY (`key`),
  KEY `Tables_gameId_fkey` (`gameId`),
  CONSTRAINT `Tables_gameId_fkey` FOREIGN KEY (`gameId`) REFERENCES `game` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `table`
--

LOCK TABLES `table` WRITE;
/*!40000 ALTER TABLE `table` DISABLE KEYS */;
INSERT INTO `table` VALUES ('782087e8-1ea6-408d-b792-75ce1b60c155','Hello',1,3,3,'2022-10-23 02:38:33.958','2022-10-23 02:38:33.958',0,'null');
/*!40000 ALTER TABLE `table` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-23 22:31:31

