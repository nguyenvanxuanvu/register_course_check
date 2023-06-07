/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `course`;
CREATE TABLE `course` (
  `id` varchar(45) NOT NULL,
  `course_name` varchar(45) NOT NULL,
  `num_credits` int NOT NULL,
  `faculty` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `course_condition`;
CREATE TABLE `course_condition` (
  `id` int NOT NULL AUTO_INCREMENT,
  `course_id` varchar(45) NOT NULL,
  `course_condition` json NOT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id_fk_idx` (`course_id`),
  CONSTRAINT `course_id_fk` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `min_max_credit`;
CREATE TABLE `min_max_credit` (
  `id` int NOT NULL AUTO_INCREMENT,
  `academic_program` varchar(45) NOT NULL,
  `semester` int NOT NULL,
  `min_credit` int NOT NULL DEFAULT '-1',
  `max_credit` int NOT NULL DEFAULT '-1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `teaching_plan`;
CREATE TABLE `teaching_plan` (
  `id` int NOT NULL AUTO_INCREMENT,
  `faculty` varchar(45) NOT NULL,
  `speciality` varchar(45) NOT NULL,
  `academic_program` varchar(45) NOT NULL,
  `semester_order` int NOT NULL,
  `course_list` longtext NOT NULL,
  `free_credit_info` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `white_list`;
CREATE TABLE `white_list` (
  `id` int NOT NULL,
  `student_id` varchar(45) NOT NULL,
  `semester` int NOT NULL,
  `min_credit` int NOT NULL,
  `max_credit` int NOT NULL,
  `description` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `course` (`id`, `course_name`, `num_credits`, `faculty`) VALUES
  ('CH1003', 'General Chemistry', 3, 'HC'), 
  ('CO1005', 'Introduction to Computing', 3, 'MT'), 
  ('CO1007', 'Discrete Structures for Computing', 4, 'MT'), 
  ('CO1023', 'Digital Systems', 3, 'MT'), 
  ('CO1027', 'Programming Fundamentals', 3, 'MT'), 
  ('CO2001', 'Professional Skills for Engineers', 3, 'MT'), 
  ('CO2003', 'Data Structures and Algorithms', 4, 'MT'), 
  ('CO2007', 'Computer Architecture', 4, 'MT'), 
  ('CO2011', 'Mathematical Modeling', 3, 'MT'), 
  ('CO2013', 'Database Systems', 4, 'MT'), 
  ('CO2017', 'Operating Systems', 3, 'MT'),
  ('CO2039', 'Advanced Programming', 3, 'MT'),
  ('CO3001', 'Software Engineering', 3, 'MT'),
  ('CO3005', 'Principles of Programming Languages', 4, 'MT'),
  ('CO3093', 'Computer Networks', 3, 'MT'),
  ('CO3335', 'Internship', 2, 'MT'),
  ('CO4029', 'Specialized Project', 2, 'MT'),
  ('CO4337', 'Capstone Project', 4, 'MT'),
  ('LA1003', 'English 1', 2, 'TN'),
  ('LA1005', 'English 2', 2, 'TN'),
  ('LA1007', 'English 3', 2, 'TN'),
  ('LA1009', 'English 4', 2, 'TN'),
  ('MI1003', 'Military Training', 0, 'GH'),
  ('MT1003', 'Calculus 1', 4, 'UD'),
  ('MT1005', 'Calculus 2', 4, 'UD'),
  ('MT1007', 'Linear Algebra', 3, 'UD'),
  ('MT2013', 'Probability and Statistics', 4, 'UD'),
  ('PH1003', 'General Physics 1', 4, 'UD'),
  ('PH1007', 'General Physics Labs', 1, 'UD'),
  ('SP1007', 'Introduction to Vietnamese Law', 2, 'UD'),
  ('SP1031', 'Marxist - Leninist Philosophy', 3, 'UD'),
  ('SP1033', 'Marxist - Leninist Political Economy', 2, 'UD'),
  ('SP1035', 'Scientific Socialism', 2, 'UD'),
  ('SP1037', 'Ho Chi Minh Ideology', 2, 'UD'),
  ('SP1039', 'History of Vietnamese Communist Party', 2, 'UD');

INSERT INTO `course_condition` (`id`, `course_id`, `course_condition`) VALUES
  (100, 'CO2003', '{\"op\":\"AND\",\"leaves\":[{\"course\":{\"type\":2,\"courseDesId\":\"CO1007\"}},{\"course\":{\"type\":1,\"courseDesId\":\"CO1027\"}}]}'),
  (101, 'CO2007', '{\"op\":\"AND\",\"leaves\":[{\"course\":{\"type\":2,\"courseDesId\":\"CO1005\"}},{\"course\":{\"type\":1,\"courseDesId\":\"CO1023\"}}]}'),
  (102, 'CO2011', '{\"course\":{\"type\":2,\"courseDesId\":\"CO1007\"}}'),
  (103, 'CO3005', '{\"course\":{\"type\":1,\"courseDesId\":\"CO2003\"}}'),
  (104, 'CO4029', '{\"course\":{\"type\":3,\"courseDesId\":\"CO3335\"}}'),
  (105, 'CO4337', '{\"op\":\"AND\",\"leaves\":[{\"course\":{\"type\":1,\"courseDesId\":\"CO3335\"}},{\"course\":{\"type\":1,\"courseDesId\":\"CO4029\"}}]}'),
  (106, 'LA1005', '{\"course\":{\"type\":1,\"courseDesId\":\"LA1003\"}}'),
  (107, 'LA1007', '{\"course\":{\"type\":1,\"courseDesId\":\"LA1005\"}}'),
  (108, 'LA1009', '{\"course\":{\"type\":1,\"courseDesId\":\"LA1007\"}}'),
  (109, 'MT1005', '{\"course\":{\"type\":2,\"courseDesId\":\"MT1003\"}}'),
  (110, 'MT2013', '{\"op\":\"AND\",\"leaves\":[{\"course\":{\"type\":2,\"courseDesId\":\"MT1003\"}},{\"course\":{\"type\":2,\"courseDesId\":\"MT1007\"}}]}'),
  (111, 'SP1033', '{\"course\":{\"type\":2,\"courseDesId\":\"SP1031\"}}'),
  (112, 'SP1035', '{\"course\":{\"type\":2,\"courseDesId\":\"SP1033\"}}'),
  (113, 'SP1037', '{\"course\":{\"type\":2,\"courseDesId\":\"SP1039\"}}'),
  (114, 'SP1039', '{\"course\":{\"type\":2,\"courseDesId\":\"SP1035\"}}');

INSERT INTO `min_max_credit` (`id`, `academic_program`, `semester`, `min_credit`, `max_credit`) VALUES
  (1, 'DT', 191, 12, 21),
  (2, 'CLC', 191, 12, 22);

INSERT INTO `teaching_plan` (`id`, `faculty`, `speciality`, `academic_program`, `semester_order`, `course_list`, `free_credit_info`) VALUES
  (1, 'KHMT', 'KHM',  'DT', 1, '[\"LA1003\",\"MI1003\",\"PE1003\",\"MT1003\",\"PH1003\",\"CO1005\",\"CO1023\"]', '[]'),
  (2, 'KHMT', 'KHM',  'DT', 2, '[\"LA1005\",\"PE1005\",\"MT1005\",\"MT1007\",\"CO1007\",\"CO1027\",\"PH1007\"]', '[]'),
  (3, 'KHMT', 'KHM',  'DT', 3, '[\"LA1007\",\"SP1031\",\"CO2007\",\"CO2011\",\"CO2003\",\"PE1007\"]', '[\n{\n\"nums\":3,\n\"group\":\"C\"\n},\n{\n\"nums\":6,\n\"group\":\"D\"\n}\n]'),
  (4, 'KHMT', 'KHM',  'DT', 4, '[\"LA1009\",\"SP1033\",\"CO2017\",\"CO2039\",\"MT2013\"]', '[]'),
  (5, 'KHMT', 'KHM',  'DT', 5, '[\"SP1035\",\"CO3093\",\"CO2013\",\"CO3001\",\"CH1003\"]', '[]'),
  (6, 'KHMT', 'KHM',  'DT', 6, '[\"SP1039\",\"CO2001\",\"CO3005\",\"CO3335\"]', '[]'),
  (7, 'KHMT', 'KHM',  'DT', 7, '[\"SP1037\",\"CO4029\"]', '[]'),
  (8, 'KHMT', 'KHM',  'DT', 8, '[\"SP1007\",\"CO4337\"]', '[]');


INSERT INTO `white_list` (`id`, `student_id`, `semester`, `min_credit`, `max_credit`, `description`) VALUES
  (1, '1915983', 191, 1, 20, 'test'),
  (2, '1915983', 192, 4, 15, NULL);