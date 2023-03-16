-- drop database if exists register_course_check;

CREATE DATABASE `register_course_check` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;


USE `register_course_check`;

CREATE TABLE `min_max_credit` (
  `int` bigint NOT NULL AUTO_INCREMENT,
  `academic_program` varchar(45) DEFAULT NULL,
  `semester` int DEFAULT NULL,
  `min_credit` int NOT NULL DEFAULT '-1',
  `max_credit` int NOT NULL DEFAULT '-1',
  PRIMARY KEY (`int`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




CREATE TABLE `course` (
  `id` varchar(45) NOT NULL,
  `course_name` varchar(45) NOT NULL,
  `num_credits` int NOT NULL,
  `faculty` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `course_condition` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `course_id` varchar(45) NOT NULL,
  `condition` json NOT NULL,
  PRIMARY KEY (`id`),
  KEY `course_id_fk_idx` (`course_id`),
  CONSTRAINT `course_id_fk` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




INSERT INTO `register_course_check`.`min_max_credit`
(`int`,`academic_program`,`semester`,`min_credit`,`max_credit`)
VALUES
('1', 'DT', '191', '5', '20');




INSERT INTO `register_course_check`.`course`
(`id`,`course_name`,`num_credits`,`faculty`)
VALUES
('CHIN', 'Chinese', '3', 'Demo'),
('ENGL', 'English', '3', 'Demo'),
('GER', 'German', '3', 'Demo'),
('SPAN', 'Spanish', '3', 'Demo'),
('ALG', 'Algebra', '3', 'Demo'),
('CALC', 'Calculus', '3', 'Demo'),
('CS', 'Computer Science', '3', 'Demo'),
('ENGR', 'Engineering', '3', 'Demo'),
('PHYS', 'Physics', '3', 'Demo'),
('CHM', 'Chemistry', '3', 'Demo'),
('BIOL', 'Biology', '3', 'Demo'),
('MBIO', 'Microbiology', '3', 'Demo'),
('PHAR', 'Pharmacy', '3', 'Demo'),
('BAND', 'Band', '3', 'Demo'),
('COM', 'Communications', '3', 'Demo'),
('ECON', 'Economics', '3', 'Demo'),
('HIST', 'History', '3', 'Demo'),
('LING', 'Linguistics', '3', 'Demo'),
('PHIL', 'Philosophy', '3', 'Demo'),
('POL', 'Political Science', '3', 'Demo'),
('PSY', 'Psychology', '3', 'Demo'),
('SOC', 'Solciology', '3', 'Demo'),
('CO1', 'AAA', '3', 'MT'),
('CO2', 'BBB', '4', 'HH'),
('CO3', 'CCC', '5', 'MT'),
('CO4', 'DDD', '3', 'CK'),
('CO5', 'New', '2', 'MT'),
('CO6', 'Test', '3', 'MT');






INSERT INTO `register_course_check`.`course_condition`
(`id`,`course_id`,`condition`)
VALUES
('10', 'CS', '{\"data\": \"ENGL-1\"}'),
('11', 'PHAR', '{\"data\":\"AND\",\"left\":{\"data\":\"CHM-1\"},\"right\":{\"data\":\"BIOL-1\"}}'),
('20', 'MBIO', '{\"data\": \"BIOL-2\"}'),
('30', 'CALC', '{\"data\": \"ALG-3\"}'),
('100', 'COM', '{\"data\":\"OR\",\"left\":{\"data\":\"ENGL-1\"},\"right\":{\"data\":\"OR\",\"left\":{\"data\":\"GER-1\"},\"right\":{\"data\":\"AND\",\"left\":{\"data\":\"SPAN-1\"},\"right\":{\"data\":\"CHIN-1\"}}}}'),
('1', 'CO1', '{\"data\": \"OR\", \"left\": {\"data\": \"CO2-1\"}, \"right\": {\"data\": \"AND\", \"left\": {\"data\": \"AND\", \"left\": {\"data\": \"CO5-1\"}, \"right\": {\"data\": \"CO6-2\"}}, \"right\": {\"data\": \"CO4-3\"}}}'),
('5', 'CO2', '{\"data\": \"CO3-1\"}'),
('6', 'CO3', '{\"data\": \"AND\", \"left\": {\"data\": \"CO4-1\"}, \"right\": {\"data\": \"CO5-1\"}}'),
('7', 'CO4', '{\"data\": \"CO2-3\"}');





