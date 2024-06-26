-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: tttn
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `username` varchar(100) NOT NULL,
  `user_password` varchar(500) NOT NULL,
  `role_id` tinyint unsigned NOT NULL,
  PRIMARY KEY (`username`),
  KEY `fk_account_role_id_idx` (`role_id`),
  CONSTRAINT `fk_account_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`role_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES ('bac@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',3),('cabcab@gmail.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('davidjohnson@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('hoang@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',3),('hoangvane@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',1),('janesmith@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('johndoe@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('phamthid@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',1);
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bill`
--

DROP TABLE IF EXISTS `bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bill` (
  `bill_id` char(25) NOT NULL,
  `t_create` datetime DEFAULT NULL,
  `tax_id` varchar(45) DEFAULT NULL,
  `tax_name` varchar(45) DEFAULT NULL,
  `customer_order_id` char(25) NOT NULL,
  `employee_id` char(25) NOT NULL,
  PRIMARY KEY (`bill_id`),
  UNIQUE KEY `customer_order_id_UNIQUE` (`customer_order_id`),
  KEY `fk_bill_customer_order_id_idx` (`customer_order_id`),
  KEY `fk_bill_employee_id_idx` (`employee_id`),
  CONSTRAINT `fk_bill_customer_order_id` FOREIGN KEY (`customer_order_id`) REFERENCES `customer_order` (`customer_order_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_bill_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bill`
--

LOCK TABLES `bill` WRITE;
/*!40000 ALTER TABLE `bill` DISABLE KEYS */;
INSERT INTO `bill` VALUES ('HO8D6dv4AfUZuv_heD2Y2','2023-08-05 16:52:52','','','h3oGekBlF16TKYMvnJ8BF','EID1689051600');
/*!40000 ALTER TABLE `bill` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `brand`
--

DROP TABLE IF EXISTS `brand`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `brand` (
  `brand_id` char(25) NOT NULL,
  `brand_name` varchar(100) NOT NULL,
  `brand_img` varchar(500) DEFAULT NULL,
  `brand_desc` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`brand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `brand`
--

LOCK TABLES `brand` WRITE;
/*!40000 ALTER TABLE `brand` DISABLE KEYS */;
INSERT INTO `brand` VALUES ('BR044023','Carlo Rossi','https://wineshop.vn/public/uploaded/product_brand/product_brand_13.jpg','A smooth and full-bodied red wine with aromas of ripe cherries, vanilla, and a touch of oak. Enjoy it with grilled meats or hearty pasta dishes.'),('BR210079','Yellow Tail','https://wineshop.vn/public/uploaded/product_brand/product_brand_15.png','A medium-bodied and well-balanced red wine with red fruit flavors, gentle tannins, and a touch of earthiness. Enjoy it with roasted poultry or grilled vegetables.'),('BR530964','Woodbridge Mondavi','https://wineshop.vn/public/uploaded/product_brand/san-mazano.jpg','A complex and aromatic white wine with tropical fruit flavors, hints of honey, and a refreshing acidity. Pair it with spicy Asian cuisine or creamy cheeses.'),('BR556729','Peter Vella','https://wineshop.vn/public/uploaded/product_brand/product_brand_16.png','A crisp and lively sparkling wine with delicate bubbles and vibrant notes of green apples and citrus. Celebrate any occasion with a glass of this sparkling delight.'),('BR689371','Barefoot Cellars','https://wineshop.vn/public/uploaded/product_brand/home-korta.jpg','An elegant and crisp white wine with refreshing citrus flavors and a touch of floral aroma. Ideal for pairing with seafood or salads.'),('BR729998','Franzia','https://wineshop.vn/public/uploaded/product_brand/francis-ford.png','A rich and velvety red wine with notes of dark berries, cocoa, and a hint of spice. Perfect for cozy evenings by the fireplace.'),('BR908082','Twin Valley','https://wineshop.vn/public/uploaded/product_brand/product_brand_9.png','A bold and robust red wine with intense blackberry and plum flavors, balanced by velvety tannins and a long, lingering finish. Perfect for steak or barbecued ribs.'),('BR964479','Sutter Home','https://wineshop.vn/public/uploaded/product_brand/product_brand_7.png','A vibrant and fruity rosé wine with flavors of fresh strawberries, watermelon, and a zesty finish. Best served chilled on warm summer days.');
/*!40000 ALTER TABLE `brand` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cart`
--

DROP TABLE IF EXISTS `cart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cart` (
  `customer_id` char(25) NOT NULL,
  `product_id` char(25) NOT NULL,
  `amount` int NOT NULL,
  PRIMARY KEY (`customer_id`,`product_id`),
  KEY `fk_cart_product_id_idx` (`product_id`),
  CONSTRAINT `fk_cart_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_cart_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart`
--

LOCK TABLES `cart` WRITE;
/*!40000 ALTER TABLE `cart` DISABLE KEYS */;
/*!40000 ALTER TABLE `cart` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `category_id` char(25) NOT NULL,
  `category_name` varchar(100) NOT NULL,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES ('TW264812','Sparkling Wines'),('TW366960','Dessert Wines'),('TW667968','Red Wines'),('TW791274','Red Wines'),('TW868383','White Wines'),('TW955788','Rosé Wines');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer` (
  `customer_id` char(25) NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `gender` tinyint(1) DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `phone_number` varchar(11) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  KEY `fk_customer_account_idx` (`email`),
  CONSTRAINT `fk_customer_account` FOREIGN KEY (`email`) REFERENCES `accounts` (`username`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES ('gj6xxjLVV8aA3SvrGJZCQ','Nguyen Dang Bac',1,'2001-07-08','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935','cabcab@gmail.com'),('hJ5Ly8WYdIrb34MJNnauQ','Nguyễn Văn A',1,'1990-01-15','Số 10, Đường ABC, Quận XYZ, Thành phố HCM','0901234567','johndoe@example.com'),('ks6DKgmOmYH5OGbWPeHg4','Trần Thị B',0,'1985-05-20','Số 20, Đường XYZ, Quận ABC, Thành phố Hà Nội','0912345678','janesmith@example.com');
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer_order`
--

DROP TABLE IF EXISTS `customer_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer_order` (
  `customer_order_id` char(25) NOT NULL,
  `t_create` datetime DEFAULT NULL,
  `full_name` varchar(100) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `phone_number` varchar(11) DEFAULT NULL,
  `t_delivery` datetime DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `payment_status` tinyint(1) DEFAULT '0',
  `employee_id` char(25) DEFAULT NULL,
  `deliverer_id` char(25) DEFAULT NULL,
  `customer_id` char(25) NOT NULL,
  PRIMARY KEY (`customer_order_id`),
  KEY `customer_order_customer_id_idx` (`customer_id`),
  KEY `fk_customer_order_employee_id_idx` (`employee_id`),
  KEY `fk_customer_order_deliverer_id_idx` (`deliverer_id`),
  CONSTRAINT `fk_customer_order_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_customer_order_deliverer_id` FOREIGN KEY (`deliverer_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_customer_order_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_order`
--

LOCK TABLES `customer_order` WRITE;
/*!40000 ALTER TABLE `customer_order` DISABLE KEYS */;
INSERT INTO `customer_order` VALUES ('3mxsj9n1rp-MwWTF0oV6p','2023-08-05 10:48:14','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935','2023-08-13 00:00:00',2,2,'EID1689051600','EID1689051620','gj6xxjLVV8aA3SvrGJZCQ'),('8GQss0qd-FNdvLnasiEVY','2023-07-18 15:26:59','Hoang Van Thu','Ho Chi Minh','0123456789','2023-07-22 00:00:00',3,2,'EID1689051600','EID1689051620','hJ5Ly8WYdIrb34MJNnauQ'),('Ap1uSGO2GTM454TDlcUT3','2023-07-20 15:26:59','Hoang Van Thu','Ho Chi Minh','0123456789','2023-07-23 00:00:00',3,2,'EID1689051600','EID1689051620','hJ5Ly8WYdIrb34MJNnauQ'),('Dp6J3Zb3NekIcEZIfaB27','2023-07-29 16:53:54','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935','2023-08-06 00:00:00',2,1,'EID1689051600','EID1689051620','gj6xxjLVV8aA3SvrGJZCQ'),('E-4la8n_b78XhPwCA8SlC','2023-08-05 10:54:04','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935',NULL,1,2,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ'),('Ewj4jd4ZwT5sy3UBy9Nur','2023-08-05 18:13:38','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935',NULL,1,2,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ'),('F3S_MJle4XxfMA9wNw2s2','2023-08-05 18:12:20','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935',NULL,1,1,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ'),('fJPIiELTjIDROGnBYUq6C','2023-07-17 15:26:59','Hoang Van Thu','Ho Chi Minh','0123456789','2023-07-25 00:00:00',3,2,'EID1689051600','EID1689051620','hJ5Ly8WYdIrb34MJNnauQ'),('h3oGekBlF16TKYMvnJ8BF','2023-07-30 20:35:24','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935','2023-08-04 00:00:00',2,2,'EID1689051600','EID1689051610','gj6xxjLVV8aA3SvrGJZCQ'),('iDgRN-DbTIl_QsbE54-0q','2023-07-24 15:26:59','Hoang Van Thu','Ho Chi Minh','0123456789','2023-07-27 00:00:00',3,2,'EID1689051600','EID1689051620','hJ5Ly8WYdIrb34MJNnauQ'),('ntKndB0rFyJy9IAnLRIms','2023-07-10 15:26:59','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935','2023-07-17 00:00:00',3,2,'EID1689051600','EID1688965200','gj6xxjLVV8aA3SvrGJZCQ'),('p1k7oHSW5F9fGyX86Mrwe','2023-07-29 21:24:01','Le Van Luon','asdasd','0982777935',NULL,1,1,NULL,NULL,'hJ5Ly8WYdIrb34MJNnauQ'),('szi3MPHMLXc1Al1Ai0fRq','2023-08-05 10:42:15','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935',NULL,1,1,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ'),('w6KwLji8g2EOY3i60Z3yc','2023-07-30 16:13:09','Nguyen Dang Bac','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0982777935',NULL,1,2,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ');
/*!40000 ALTER TABLE `customer_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer_order_detail`
--

DROP TABLE IF EXISTS `customer_order_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer_order_detail` (
  `customer_order_detail_id` char(25) NOT NULL,
  `product_id` char(25) NOT NULL,
  `customer_order_id` char(25) NOT NULL,
  `amount` int DEFAULT NULL,
  `cost` float DEFAULT NULL,
  PRIMARY KEY (`customer_order_detail_id`),
  KEY `fk_customer_order_detail_product_id_idx` (`product_id`),
  KEY `fk_customer_order_detail_customer_order_id_idx` (`customer_order_id`),
  CONSTRAINT `fk_customer_order_detail_customer_order_id` FOREIGN KEY (`customer_order_id`) REFERENCES `customer_order` (`customer_order_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_customer_order_detail_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_order_detail`
--

LOCK TABLES `customer_order_detail` WRITE;
/*!40000 ALTER TABLE `customer_order_detail` DISABLE KEYS */;
INSERT INTO `customer_order_detail` VALUES ('0PBFkPOpEsf6n7TkLnei4','jubCMIqrRxcXAuS0W776J','h3oGekBlF16TKYMvnJ8BF',5,50),('2I8RGj-JC-wrlzA32ndIP','TRs0bA-26BsBKBNPKbp32','Ewj4jd4ZwT5sy3UBy9Nur',1,50),('3K3OaMHMoKBn7ntGha-LY','jubCMIqrRxcXAuS0W776J','E-4la8n_b78XhPwCA8SlC',1,50),('629IBoS10PyrA8IJIvyHt','fK34Dy-iitHdBj0U-36W4','ntKndB0rFyJy9IAnLRIms',1,50),('6diev8TVe_vLIJR34XRV3','jubCMIqrRxcXAuS0W776J','3mxsj9n1rp-MwWTF0oV6p',1,50),('8HdzMPc3oni3gSidL-Nvo','fK34Dy-iitHdBj0U-36W4','szi3MPHMLXc1Al1Ai0fRq',1,50),('aYQaPDqAfqRK9E6crFvX6','jubCMIqrRxcXAuS0W776J','Dp6J3Zb3NekIcEZIfaB27',1,50),('cE6r8g9iG3EnRIGJQMNDw','jubCMIqrRxcXAuS0W776J','iDgRN-DbTIl_QsbE54-0q',1,49),('f18VB-rtEN0N2uDkI8lAQ','NYDrjSfRAcpyG8SNNN7kW','Ap1uSGO2GTM454TDlcUT3',2,49),('HTFfGxLPQBvSCSVtQnUqa','7wS-8EN5-KmDMSumetyIS','ntKndB0rFyJy9IAnLRIms',2,24.5),('MYPHxNJ3poMSkIFHpiGPU','jubCMIqrRxcXAuS0W776J','w6KwLji8g2EOY3i60Z3yc',1,50),('NE8J9O0z9C_AmbWgq5mV7','gK4WEYrHboLOrD6l-GWPA','h3oGekBlF16TKYMvnJ8BF',1,50),('nSOpysEmzN-No2jJeW6Si','NYDrjSfRAcpyG8SNNN7kW','p1k7oHSW5F9fGyX86Mrwe',2,49),('O1NmgMgko_1WXgzILpISn','TRs0bA-26BsBKBNPKbp32','h3oGekBlF16TKYMvnJ8BF',5,50),('OahR1YyTC1kmDbXRNbGXY','gK4WEYrHboLOrD6l-GWPA','ntKndB0rFyJy9IAnLRIms',1,50),('OHmxYhEhqv4_MJ_NBzITw','TRs0bA-26BsBKBNPKbp32','8GQss0qd-FNdvLnasiEVY',2,49),('UNWp1RInRp6VI65gBJINK','fK34Dy-iitHdBj0U-36W4','Dp6J3Zb3NekIcEZIfaB27',1,50),('VXIaeIv_oaiF-NLr304gN','TRs0bA-26BsBKBNPKbp32','F3S_MJle4XxfMA9wNw2s2',1,50),('Ws-JMze0F8OvOXHhWNQA0','7wS-8EN5-KmDMSumetyIS','h3oGekBlF16TKYMvnJ8BF',1,24.5),('Y_F3YOFaXNBHFGsduG4QZ','jubCMIqrRxcXAuS0W776J','fJPIiELTjIDROGnBYUq6C',3,49);
/*!40000 ALTER TABLE `customer_order_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employee`
--

DROP TABLE IF EXISTS `employee`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employee` (
  `employee_id` char(25) NOT NULL,
  `full_name` varchar(100) DEFAULT NULL,
  `gender` tinyint(1) DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `phone_number` varchar(11) DEFAULT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`employee_id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  KEY `email_idx` (`email`),
  CONSTRAINT `fk_employee_account` FOREIGN KEY (`email`) REFERENCES `accounts` (`username`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employee`
--

LOCK TABLES `employee` WRITE;
/*!40000 ALTER TABLE `employee` DISABLE KEYS */;
INSERT INTO `employee` VALUES ('EID1688965200','Hoàng Văn E',1,'1988-09-25','Số 30, Đường STU, Quận VWX, Thành phố Hải Phòng','0945678901','hoangvane@example.com'),('EID1689051600','Phạm Thị D',0,'1992-03-12','Số 15, Đường MNO, Quận PQR, Thành phố Cần Thơ','0934567890','phamthid@example.com'),('EID1689051610','Hoàng Shipper',1,'1992-03-12','Số 15, Đường MNO, Quận PQR, Thành phố Cần Thơ','0934567890','hoang@example.com'),('EID1689051620','Bắc Shipper',1,'1992-03-12','Số 15, Đường MNO, Quận PQR, Thành phố Cần Thơ','0934567890','bac@example.com');
/*!40000 ALTER TABLE `employee` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `order_id` char(25) NOT NULL,
  `t_create` datetime DEFAULT NULL,
  `provider_id` char(25) NOT NULL,
  `employee_id` char(25) NOT NULL,
  PRIMARY KEY (`order_id`),
  KEY `fk_order_provider_id_idx` (`provider_id`),
  KEY `fk_order_employee_id_idx` (`employee_id`),
  CONSTRAINT `fk_order_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_order_provider_id` FOREIGN KEY (`provider_id`) REFERENCES `provider` (`provider_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_detail`
--

DROP TABLE IF EXISTS `order_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_detail` (
  `order_id` char(25) NOT NULL,
  `product_id` char(25) NOT NULL,
  `amount` int DEFAULT NULL,
  `cost` float DEFAULT NULL,
  PRIMARY KEY (`order_id`,`product_id`),
  KEY `fk_product_order_id_idx` (`product_id`),
  CONSTRAINT `fk_order_product_id` FOREIGN KEY (`order_id`) REFERENCES `order` (`order_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_product_order_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_detail`
--

LOCK TABLES `order_detail` WRITE;
/*!40000 ALTER TABLE `order_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `product_id` char(25) NOT NULL,
  `product_name` varchar(100) NOT NULL,
  `cost` float DEFAULT NULL,
  `product_img` varchar(500) DEFAULT NULL,
  `description` varchar(500) DEFAULT NULL,
  `inventory_number` int DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `is_new` tinyint(1) DEFAULT '0',
  `brand_id` char(25) NOT NULL,
  `category_id` char(25) NOT NULL,
  PRIMARY KEY (`product_id`),
  KEY `category_id_idx` (`category_id`),
  KEY `brand_id_idx` (`brand_id`),
  CONSTRAINT `fk_product_brand_id` FOREIGN KEY (`brand_id`) REFERENCES `brand` (`brand_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_product_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES ('7wS-8EN5-KmDMSumetyIS','Korta Reseva De Familia',49,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',0,'Stocking',1,'BR530964','TW667968'),('fK34Dy-iitHdBj0U-36W4','Airén',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/071102-1.4913eea891cf816624c4a3b6dfad4652.jpg','Rosé wines are represented in the best way with this Bardolino Chiaretto. By blending several red varieties, a unique, elegant and fragrant rosé is produced. Along with the traditional fruity aromas, this wine also displays a unique thyme note towards the end.',16,'Stocking',0,'BR210079','TW366960'),('fZDcOL12Au_UBjqAv6kHj','Wine Isla de Maipo, Cabernet Sauvignon, 2018',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092104-1.1fdcf7cee862d06e0d6917c56993a1d1.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',1,'BR210079','TW366960'),('gK4WEYrHboLOrD6l-GWPA','Port Quinta das Arcas, Palmira Tawny Port',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/brut-roz-cricova.a2b1f0ea8366b7db9142bc8bceefcbc0.jpg','As the rosé capital of the world, you can only expect the very best from this Provence rosé. It is a complex and expressive wine that offers both vibrant fruit and floral notes. It is a true example of how exhilarating a rosé can be.',14,'Stocking',1,'BR530964','TW667968'),('jubCMIqrRxcXAuS0W776J','Wine Altos Las Hormigas, Malbec Terroir Luján De Cuyo, 2016',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/052201-11.d79ca6ae9c38bb0c198100a8f4a50100.jpg','If ever there was a moreish rosé, it is this award-winning, aromatic wine. Not only does the wine offer vibrant flavors and acidity, but also offers layers of intense citrus and stone-fruit notes that leave you wanting more.',7,'Stocking',0,'BR044023','TW264812'),('mhLO-C5PXeTEIFT24fJDm','Frizzante Wine Arione, Bonarda dell\'Oltrepo Pavese 2021',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/101206-1.5707e31f0c1da611c0335450bcc75152.jpg','This rosé was produced using the utmost delicate practices, which has added immense to its uniqueness. By whole-bunch basket pressing, this wine expresses vibrancy, freshness and clean summer fruits like no other.',20,'Stocking',1,'BR210079','TW366960'),('NYDrjSfRAcpyG8SNNN7kW','Sauvignon Blanc',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/052203-111.9a6e82e23ea510e477eff5a959f5f414.jpg','This unique rosé combines two unique Portuguese varieties: Espadeiro and Touriga Nacional, which are hand-picked and pressed to create this slightly fizzy wine. It is an exceptional example of unique Portuguese wines.',16,'Stocking',1,'BR530964','TW667968'),('TRs0bA-26BsBKBNPKbp32','Tempranillo',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/072502-1.d88f1d989f0c6c216e4e7e6e7d32fba1.jpg','There is a charm that comes with a pink sparkling wine; especially when quality-winemaking is involved. This Pinot Noir sparkling wine has the perfect balance of structure, fruit, tannin and refreshing acidity that is suited for any celebratory occasion.',11,'Stocking',0,'BR210079','TW366960');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotion`
--

DROP TABLE IF EXISTS `promotion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `promotion` (
  `promotion_id` char(25) NOT NULL,
  `promotion_name` varchar(100) DEFAULT NULL,
  `date_start` date DEFAULT NULL,
  `date_end` date DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  `employee_id` char(25) NOT NULL,
  PRIMARY KEY (`promotion_id`),
  KEY `employee_id_idx` (`employee_id`),
  CONSTRAINT `fk_promotion_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotion`
--

LOCK TABLES `promotion` WRITE;
/*!40000 ALTER TABLE `promotion` DISABLE KEYS */;
INSERT INTO `promotion` VALUES ('N5sMhJnFE_3GQXHV9UTUa','new promotion','2023-07-11','2023-08-22','empty','EID1689051600');
/*!40000 ALTER TABLE `promotion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `promotion_detail`
--

DROP TABLE IF EXISTS `promotion_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `promotion_detail` (
  `product_id` char(25) NOT NULL,
  `promotion_id` char(25) NOT NULL,
  `discount_percentage` float DEFAULT NULL,
  PRIMARY KEY (`product_id`,`promotion_id`),
  KEY `fk_product_promotion_id_idx` (`promotion_id`),
  CONSTRAINT `fk_product_promotion_id` FOREIGN KEY (`promotion_id`) REFERENCES `promotion` (`promotion_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_promotion_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `promotion_detail`
--

LOCK TABLES `promotion_detail` WRITE;
/*!40000 ALTER TABLE `promotion_detail` DISABLE KEYS */;
INSERT INTO `promotion_detail` VALUES ('7wS-8EN5-KmDMSumetyIS','N5sMhJnFE_3GQXHV9UTUa',0.5),('mhLO-C5PXeTEIFT24fJDm','N5sMhJnFE_3GQXHV9UTUa',0.2);
/*!40000 ALTER TABLE `promotion_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `provide_product`
--

DROP TABLE IF EXISTS `provide_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `provide_product` (
  `provider_id` char(25) NOT NULL,
  `product_id` char(25) NOT NULL,
  PRIMARY KEY (`provider_id`,`product_id`),
  KEY `fk_product_id_idx` (`product_id`),
  CONSTRAINT `fk_product_provider_id` FOREIGN KEY (`provider_id`) REFERENCES `provider` (`provider_id`),
  CONSTRAINT `fk_provider_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `provide_product`
--

LOCK TABLES `provide_product` WRITE;
/*!40000 ALTER TABLE `provide_product` DISABLE KEYS */;
/*!40000 ALTER TABLE `provide_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `provider`
--

DROP TABLE IF EXISTS `provider`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `provider` (
  `provider_id` char(25) NOT NULL,
  `provider_name` varchar(100) NOT NULL,
  `address` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`provider_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `provider`
--

LOCK TABLES `provider` WRITE;
/*!40000 ALTER TABLE `provider` DISABLE KEYS */;
INSERT INTO `provider` VALUES ('SUP001','Công ty Thực phẩm Minh Trang','Số 123, Đường Trần Phú, Quận 1, Thành phố Hồ Chí Minh','minhtrang@example.com'),('SUP002','Nhà cung cấp Hùng Vương','Số 456, Đường Hùng Vương, Quận 5, Thành phố Hải Phòng','hungvuong@example.com'),('SUP003','Cửa hàng Tấn Phát','Số 789, Đường Nguyễn Văn Linh, Quận 7, Thành phố Đà Nẵng','tanphat@example.com'),('SUP004','Nhà cung cấp Ánh Dương','Số 321, Đường Lê Lợi, Quận Bình Thạnh, Thành phố Cần Thơ','anhduong@example.com'),('SUP005','Công ty Thực phẩm Xuân Hương','Số 567, Đường Lê Duẩn, Quận 3, Thành phố Hà Nội','xuanhuong@example.com');
/*!40000 ALTER TABLE `provider` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `receipt`
--

DROP TABLE IF EXISTS `receipt`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `receipt` (
  `receipt_id` char(25) NOT NULL,
  `t_create` datetime DEFAULT NULL,
  `employee_id` char(25) NOT NULL,
  `order_id` char(25) NOT NULL,
  PRIMARY KEY (`receipt_id`),
  UNIQUE KEY `order_id_UNIQUE` (`order_id`),
  KEY `fk_receipt_employee_id_idx` (`employee_id`),
  KEY `fk_receipt_order_id_idx` (`order_id`),
  CONSTRAINT `fk_receipt_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_receipt_order_id` FOREIGN KEY (`order_id`) REFERENCES `order` (`order_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `receipt`
--

LOCK TABLES `receipt` WRITE;
/*!40000 ALTER TABLE `receipt` DISABLE KEYS */;
/*!40000 ALTER TABLE `receipt` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `receipt_detail`
--

DROP TABLE IF EXISTS `receipt_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `receipt_detail` (
  `receipt_id` char(25) NOT NULL,
  `product_id` char(25) NOT NULL,
  `amount` int DEFAULT NULL,
  `cost` float DEFAULT NULL,
  PRIMARY KEY (`receipt_id`,`product_id`),
  KEY `fk_receipt_product_id_idx` (`product_id`),
  CONSTRAINT `fk_product_receipt_id` FOREIGN KEY (`receipt_id`) REFERENCES `receipt` (`receipt_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_receipt_product_id` FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `receipt_detail`
--

LOCK TABLES `receipt_detail` WRITE;
/*!40000 ALTER TABLE `receipt_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `receipt_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `return_order`
--

DROP TABLE IF EXISTS `return_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `return_order` (
  `return_order_id` char(25) NOT NULL,
  `t_create` datetime DEFAULT NULL,
  `customer_order_id` char(25) NOT NULL,
  `employee_id` char(25) NOT NULL,
  PRIMARY KEY (`return_order_id`),
  KEY `fk_return_order_customer_order_id_idx` (`customer_order_id`),
  KEY `fk_return_order_employee_id_idx` (`employee_id`),
  CONSTRAINT `fk_return_order_customer_order_id` FOREIGN KEY (`customer_order_id`) REFERENCES `customer_order` (`customer_order_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_return_order_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `return_order`
--

LOCK TABLES `return_order` WRITE;
/*!40000 ALTER TABLE `return_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `return_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `return_order_detail`
--

DROP TABLE IF EXISTS `return_order_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `return_order_detail` (
  `return_order_id` char(25) NOT NULL,
  `customer_order_detail_id` char(25) NOT NULL,
  `amount` int DEFAULT NULL,
  PRIMARY KEY (`return_order_id`,`customer_order_detail_id`),
  KEY `fk_return_order_detail_customer_order_detail_id_idx` (`customer_order_detail_id`),
  CONSTRAINT `fk_return_order_detail_customer_order_detail_id` FOREIGN KEY (`customer_order_detail_id`) REFERENCES `customer_order_detail` (`customer_order_detail_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_return_order_detail_return_order_id` FOREIGN KEY (`return_order_id`) REFERENCES `return_order` (`return_order_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `return_order_detail`
--

LOCK TABLES `return_order_detail` WRITE;
/*!40000 ALTER TABLE `return_order_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `return_order_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `role_id` tinyint unsigned NOT NULL,
  `role_name` varchar(100) NOT NULL,
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'admin'),(2,'client'),(3,'deliverer ');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'tttn'
--
/*!50003 DROP PROCEDURE IF EXISTS `CalculateRevenueByDateRange` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = '' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `CalculateRevenueByDateRange`(
  IN fromDate DATE,
  IN toDate DATE
)
BEGIN
  SELECT t_create as `date`, count(co.customer_order_id) as total_amount, (cod.amount*cod.cost) as revenue
  FROM customer_order as co
  INNER JOIN customer_order_detail as cod ON cod.customer_order_id= co.customer_order_id
  WHERE co.t_create >= fromDate AND co.t_create <= toDate AND co.status = 3 
	AND co.customer_order_id NOT IN 
		(SELECT customer_order_id from return_order
		WHERE t_create >= fromDate AND t_create <= toDate)
  GROUP BY t_create
  ORDER BY t_create;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-06 23:32:12
