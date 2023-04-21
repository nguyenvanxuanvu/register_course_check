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
  `id` int NOT NULL,
  `course_id` varchar(45) NOT NULL,
  `condition` json NOT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id_fk_idx` (`course_id`),
  CONSTRAINT `course_id_fk` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
('ALG', 'Algebra', 3, 'Demo');
INSERT INTO `course` (`id`, `course_name`, `num_credits`, `faculty`) VALUES
('BAND', 'Band', 3, 'Demo');
INSERT INTO `course` (`id`, `course_name`, `num_credits`, `faculty`) VALUES
('BIOL', 'Biology', 3, 'Demo');
INSERT INTO `course` (`id`, `course_name`, `num_credits`, `faculty`) VALUES
('C S', 'Computer Science', 3, 'Demo'),
('CALC', 'Calculus', 3, 'Demo'),
('CHIN', 'Chinese', 3, 'Demo'),
('CHM', 'Chemistry', 3, 'Demo'),
('CO1', 'AAA', 3, 'MT'),
('CO2', 'BBB', 4, 'HH'),
('CO3', 'CCC', 5, 'MT'),
('CO4', 'DDD', 3, 'CK'),
('CO5', 'New', 2, 'MT'),
('CO6', 'Test', 3, 'MT'),
('COM', 'Communications', 3, 'Demo'),
('ECON', 'Economics', 3, 'Demo'),
('ENGL', 'English', 3, 'Demo'),
('ENGR', 'Engineering', 3, 'Demo'),
('GER', 'German', 3, 'Demo'),
('HIST', 'History', 3, 'Demo'),
('LING', 'Linguistics', 3, 'Demo'),
('MBIO', 'Microbiology', 3, 'Demo'),
('PHAR', 'Pharmacy', 3, 'Demo'),
('PHIL', 'Philosophy', 3, 'Demo'),
('PHYS', 'Physics', 3, 'Demo'),
('POL', 'Political Science', 3, 'Demo'),
('PSY', 'Psychology', 3, 'Demo'),
('SOC', 'Solciology', 3, 'Demo'),
('SPAN', 'Spanish', 3, 'Demo'),
('test', 'test', 3, 'Demo');

INSERT INTO `course_condition` (`id`, `course_id`, `condition`) VALUES
(105, 'C S', '{\"course\": {\"type\": 1, \"courseDesId\": \"ENGL\"}}');
INSERT INTO `course_condition` (`id`, `course_id`, `condition`) VALUES
(106, 'ENGR', '{\"course\": {\"type\": 1, \"courseDesId\": \"GER\"}}');
INSERT INTO `course_condition` (`id`, `course_id`, `condition`) VALUES
(107, 'PHYS', '{\"course\": {\"type\": 1, \"courseDesId\": \"SPAN\"}}');
INSERT INTO `course_condition` (`id`, `course_id`, `condition`) VALUES
(108, 'BIOL', '{\"course\": {\"type\": 2, \"courseDesId\": \"CHM\"}}'),
(109, 'MBIO', '{\"course\": {\"type\": 2, \"courseDesId\": \"BIOL\"}}'),
(110, 'ALG', '{\"course\": {\"type\": 3, \"courseDesId\": \"CALC\"}}'),
(111, 'CALC', '{\"course\": {\"type\": 3, \"courseDesId\": \"ALG\"}}'),
(112, 'COM', '{\"op\": \"OR\", \"leaves\": [{\"course\": {\"type\": 1, \"courseDesId\": \"ENGL\"}}, {\"course\": {\"type\": 1, \"courseDesId\": \"GER\"}}, {\"op\": \"AND\", \"leaves\": [{\"course\": {\"type\": 1, \"courseDesId\": \"SPAN\"}}, {\"course\": {\"type\": 1, \"courseDesId\": \"CHIN\"}}]}]}'),
(113, 'PHAR', '{\"op\": \"AND\", \"leaves\": [{\"course\": {\"type\": 1, \"courseDesId\": \"CHM\"}}, {\"course\": {\"type\": 1, \"courseDesId\": \"BIOL\"}}]}'),
(114, 'test', '{\"op\": \"OR\", \"leaves\": [{\"course\": {\"type\": 1, \"courseDesId\": \"ENGL\"}}, {\"course\": {\"type\": 1, \"courseDesId\": \"GER\"}}, {\"op\": \"AND\", \"leaves\": [{\"course\": {\"type\": 1, \"courseDesId\": \"SPAN\"}}, {\"course\": {\"type\": 1, \"courseDesId\": \"CHIN\"}}]}]}');

INSERT INTO `min_max_credit` (`id`, `academic_program`, `semester`, `min_credit`, `max_credit`) VALUES
(1, 'DT', 191, 14, 21);
INSERT INTO `min_max_credit` (`id`, `academic_program`, `semester`, `min_credit`, `max_credit`) VALUES
(2, 'CLC', 191, 12, 22);


INSERT INTO `teaching_plan` (`id`, `faculty`, `speciality`, `academic_program`, `semester_order`, `course_list`, `free_credit_info`) VALUES
(1, 'KHMT', 'KHM', 'DT', 3, '[\"C S\", \"CALC\"]', '[\n  {\n    \"nums\": 3,\n    \"group\": \"C\"\n  },\n  {\n    \"nums\": 6,\n    \"group\": \"D\"\n  }\n]');


INSERT INTO `white_list` (`id`, `student_id`, `semester`, `min_credit`, `max_credit`, `description`) VALUES
(1, '1915983', 191, 1, 20, 'test');
INSERT INTO `white_list` (`id`, `student_id`, `semester`, `min_credit`, `max_credit`, `description`) VALUES
(2, '1915983', 192, 4, 15, NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;