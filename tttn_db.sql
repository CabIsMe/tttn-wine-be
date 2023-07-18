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
INSERT INTO `accounts` VALUES ('cabcab@gmail.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('davidjohnson@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('hoangvane@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',1),('janesmith@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('johndoe@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',2),('phamthid@example.com','$2a$10$/oULmiYRaK7.Mcb7CS5.VeNxE54nYMjgo7d9pgFuzOnIwNKj.fsGe',1);
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
INSERT INTO `customer` VALUES ('gj6xxjLVV8aA3SvrGJZCQ','Lê Quang C',1,'1995-07-10','Số 5, Đường DEF, Quận GHI, Thành phố Đà Nẵng','0923456789','cabcab@gmail.com'),('hJ5Ly8WYdIrb34MJNnauQ','Nguyễn Văn A',1,'1990-01-15','Số 10, Đường ABC, Quận XYZ, Thành phố HCM','0901234567','johndoe@example.com'),('ks6DKgmOmYH5OGbWPeHg4','Trần Thị B',0,'1985-05-20','Số 20, Đường XYZ, Quận ABC, Thành phố Hà Nội','0912345678','janesmith@example.com');
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
INSERT INTO `customer_order` VALUES ('lB7EGZNOdKOIx-CmRMS9_','2023-07-17 11:59:58','','aslkjdas','0238478','2023-07-19 00:00:00',1,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ'),('M4HCLARFh9x92nqKn8HSA','2023-07-17 12:00:35','','aslkjdas','0238478','2023-07-19 00:00:00',1,NULL,NULL,'gj6xxjLVV8aA3SvrGJZCQ');
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
INSERT INTO `customer_order_detail` VALUES ('2AAQQaD2hFBlYm9QItHeG','7wS-8EN5-KmDMSumetyIS','M4HCLARFh9x92nqKn8HSA',5,100),('oEGbtgJQ9vOiTD984UKk7','7wS-8EN5-KmDMSumetyIS','lB7EGZNOdKOIx-CmRMS9_',4,100);
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
INSERT INTO `employee` VALUES ('EID1688965200','Hoàng Văn E',1,'1988-09-25','Số 30, Đường STU, Quận VWX, Thành phố Hải Phòng','0945678901','hoangvane@example.com'),('EID1689051600','Phạm Thị D',0,'1992-03-12','Số 15, Đường MNO, Quận PQR, Thành phố Cần Thơ','0934567890','phamthid@example.com');
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
INSERT INTO `product` VALUES ('7wS-8EN5-KmDMSumetyIS','Korta Reseva De Familia',49,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',7,'Stocking',1,'BR530964','TW667968'),('fK34Dy-iitHdBj0U-36W4','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Out of stock',0,'BR210079','TW366960'),('fZDcOL12Au_UBjqAv6kHj','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',1,'BR210079','TW366960'),('gK4WEYrHboLOrD6l-GWPA','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',1,'BR530964','TW667968'),('jubCMIqrRxcXAuS0W776J','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',0,'BR044023','TW264812'),('mhLO-C5PXeTEIFT24fJDm','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',1,'BR210079','TW366960'),('NYDrjSfRAcpyG8SNNN7kW','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',1,'BR530964','TW667968'),('TRs0bA-26BsBKBNPKbp32','Korta Reseva De Familia',50,'https://vinoteka.vn/assets/components/phpthumbof/cache/092121-1.1c7d8cfea75f219576db460999053e55.jpg','This prestigious Bordeaux wine boasts a deep red color with rich aromas of black currants, cedar, and tobacco. On the palate, it offers a harmonious blend of dark fruits, velvety tannins, and a long, elegant finish. Perfect for special occasions and pairing with fine cuisine.',20,'Stocking',0,'BR210079','TW366960');
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
INSERT INTO `promotion` VALUES ('N5sMhJnFE_3GQXHV9UTUa','new promotion','2023-07-11','2023-07-22','empty','EID1689051600');
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
INSERT INTO `roles` VALUES (1,'admin'),(2,'client');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-18 20:35:19
