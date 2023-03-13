-- -------------------------------------------------------------
-- TablePlus 5.0.2(458)
--
-- https://tableplus.com/
--
-- Database: register_course_check
-- Generation Time: 2022-12-09 18:16:30.0820
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `min_max_credit` (
  `int` bigint NOT NULL AUTO_INCREMENT,
  `academic_program` varchar(45) DEFAULT NULL,
  `semester` int DEFAULT NULL,
  `min_credit` int DEFAULT '0',
  `max_credit` int DEFAULT '0',
  PRIMARY KEY (`int`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `subject` (
  `id` varchar(45) NOT NULL,
  `subject_name` varchar(45) NOT NULL,
  `num_credits` int NOT NULL,
  `faculty` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `subject_condition` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `subject_id` varchar(45) NOT NULL,
  `condition` json NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO `register_course_check`.`min_max_credit`
(`int`,`academic_program`,`semester`,`min_credit`,`max_credit`)
VALUES
('1', 'DT', '191', '1', '10');



INSERT INTO `register_course_check`.`subject`
(`id`,`subject_name`,`num_credits`,`faculty`)
VALUES
('CO1', 'AAA', '3', 'MT'),
('CO2', 'BBB', '4', 'HH'),
('CO3', 'CCC', '5', 'MT'),
('CO4', 'DDD', '3', 'CK');




INSERT INTO `register_course_check`.`subject_condition`
(`id`,`subject_id`,`condition`)
VALUES
('1', 'CO1', '{\"data\": \"OR\", \"left\": {\"data\": \"CO2-1\"}, \"right\": {\"data\": \"AND\", \"left\": {\"data\": \"AND\", \"left\": {\"data\": \"CO5-1\"}, \"right\": {\"data\": \"CO6-2\"}}, \"right\": {\"data\": \"CO4-3\"}}}'),
('2', 'CO2', '{\"data\": \"CO3-1\"}');




