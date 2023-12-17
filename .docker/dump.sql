-- MySQL dump 10.13  Distrib 8.2.0, for macos14.0 (arm64)
--
-- Host: 127.0.0.1    Database: contentdb
-- ------------------------------------------------------
-- Server version	8.2.0

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
-- Table structure for table `incident_types`
--

DROP TABLE IF EXISTS `incident_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `incident_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incident_types`
--

/*!40000 ALTER TABLE `incident_types` DISABLE KEYS */;
INSERT INTO `incident_types` VALUES (1,'CSIRTaaS'),(2,'RapidResponse');
/*!40000 ALTER TABLE `incident_types` ENABLE KEYS */;

--
-- Table structure for table `incidents`
--

DROP TABLE IF EXISTS `incidents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `incidents` (
  `id` varchar(50) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `severity` enum('Low','Medium','High','Critical') DEFAULT 'Low',
  `status` enum('Pending','Open','Active','Closed') DEFAULT 'Pending',
  `type` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type` (`type`),
  CONSTRAINT `incidents_ibfk_1` FOREIGN KEY (`type`) REFERENCES `incident_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incidents`
--

/*!40000 ALTER TABLE `incidents` DISABLE KEYS */;
INSERT INTO `incidents` VALUES ('b595bee6-d8c7-4b3b-b071-1e45c2103002','RapidResponse Case 1','Low','Pending',1),('b595bee6-d8c7-4b3b-b072-1e45c2103002','RapidResponse Case 2','Critical','Pending',1);
/*!40000 ALTER TABLE `incidents` ENABLE KEYS */;

--
-- Table structure for table `ioc_types`
--

DROP TABLE IF EXISTS `ioc_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ioc_types` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ioc_types`
--

/*!40000 ALTER TABLE `ioc_types` DISABLE KEYS */;
INSERT INTO `ioc_types` VALUES (1,'URL'),(2,'DOMAIN');
/*!40000 ALTER TABLE `ioc_types` ENABLE KEYS */;

--
-- Table structure for table `iocs`
--

DROP TABLE IF EXISTS `iocs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `iocs` (
  `id` varchar(50) NOT NULL,
  `value` varchar(50) NOT NULL,
  `iocType` int NOT NULL,
  `verdict` enum('Neutral','Benigne','Malicious') DEFAULT 'Neutral',
  PRIMARY KEY (`id`),
  KEY `iocType` (`iocType`),
  CONSTRAINT `iocs_ibfk_1` FOREIGN KEY (`iocType`) REFERENCES `ioc_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `iocs`
--

/*!40000 ALTER TABLE `iocs` DISABLE KEYS */;
INSERT INTO `iocs` VALUES ('49c9793c-8492-468b-8ae0-64e37eb01fa0','youtube.com',2,'Neutral'),('b7a8ae7e-55ae-4983-bd36-ba26c5320487','google.com',2,'Neutral');
/*!40000 ALTER TABLE `iocs` ENABLE KEYS */;

--
-- Table structure for table `iocs_incidents`
--

DROP TABLE IF EXISTS `iocs_incidents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `iocs_incidents` (
  `id` int NOT NULL AUTO_INCREMENT,
  `iocId` varchar(50) DEFAULT NULL,
  `incidentId` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `iocId` (`iocId`),
  KEY `incidentId` (`incidentId`),
  CONSTRAINT `iocs_incidents_ibfk_1` FOREIGN KEY (`iocId`) REFERENCES `iocs` (`id`),
  CONSTRAINT `iocs_incidents_ibfk_2` FOREIGN KEY (`incidentId`) REFERENCES `incidents` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `iocs_incidents`
--

/*!40000 ALTER TABLE `iocs_incidents` DISABLE KEYS */;
/*!40000 ALTER TABLE `iocs_incidents` ENABLE KEYS */;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task` (
  `id` varchar(255) NOT NULL,
  `Title` varchar(255) NOT NULL,
  `Description` text NOT NULL,
  `Assignee` varchar(255) DEFAULT NULL,
  `Incident` varchar(255) DEFAULT NULL,
  `Status` enum('Open','In Progress','Done') NOT NULL DEFAULT 'Open',
  `Priority` enum('Low','Medium','High','Critical') DEFAULT 'Low',
  PRIMARY KEY (`id`),
  KEY `Assignee` (`Assignee`),
  KEY `Incident` (`Incident`),
  CONSTRAINT `task_ibfk_1` FOREIGN KEY (`Assignee`) REFERENCES `users` (`id`),
  CONSTRAINT `task_ibfk_2` FOREIGN KEY (`Incident`) REFERENCES `incidents` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task`
--

/*!40000 ALTER TABLE `task` DISABLE KEYS */;
/*!40000 ALTER TABLE `task` ENABLE KEYS */;

--
-- Table structure for table `task_comments`
--

DROP TABLE IF EXISTS `task_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task_comments` (
  `id` varchar(255) NOT NULL COMMENT 'UUID',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'Create Time',
  `content` text NOT NULL,
  `writer` varchar(255) NOT NULL,
  `task` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `writer` (`writer`),
  KEY `task` (`task`),
  CONSTRAINT `task_comments_ibfk_1` FOREIGN KEY (`writer`) REFERENCES `users` (`id`),
  CONSTRAINT `task_comments_ibfk_2` FOREIGN KEY (`task`) REFERENCES `task` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task_comments`
--

/*!40000 ALTER TABLE `task_comments` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_comments` ENABLE KEYS */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(50) NOT NULL,
  `firstName` varchar(255) NOT NULL,
  `lastName` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_At` varchar(50) DEFAULT NULL,
  `fullname` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

--
-- Table structure for table `worklogs`
--

DROP TABLE IF EXISTS `worklogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `worklogs` (
  `id` varchar(50) NOT NULL,
  `writerId` varchar(50) NOT NULL,
  `incidentId` varchar(50) NOT NULL,
  `content` text NOT NULL,
  `created_at` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `writerId` (`writerId`),
  KEY `incidentId` (`incidentId`),
  CONSTRAINT `worklogs_ibfk_1` FOREIGN KEY (`writerId`) REFERENCES `users` (`id`),
  CONSTRAINT `worklogs_ibfk_2` FOREIGN KEY (`incidentId`) REFERENCES `incidents` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `worklogs`
--

/*!40000 ALTER TABLE `worklogs` DISABLE KEYS */;
/*!40000 ALTER TABLE `worklogs` ENABLE KEYS */;

--
-- Dumping routines for database 'contentdb'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-17 19:56:34
