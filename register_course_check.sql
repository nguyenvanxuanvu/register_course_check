CREATE DATABASE `register_course_check` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;



CREATE TABLE `min_max_credit` (
  `int` bigint NOT NULL AUTO_INCREMENT,
  `academic_program` varchar(45) DEFAULT NULL,
  `semester` int DEFAULT NULL,
  `min_credit` int NOT NULL DEFAULT '-1',
  `max_credit` int NOT NULL DEFAULT '-1',
  PRIMARY KEY (`int`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




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
  PRIMARY KEY (`id`),
  KEY `subject_id_fk_idx` (`subject_id`),
  CONSTRAINT `subject_id_fk` FOREIGN KEY (`subject_id`) REFERENCES `subject` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




INSERT INTO `register_course_check`.`min_max_credit`
(`int`,`academic_program`,`semester`,`min_credit`,`max_credit`)
VALUES
('1', 'DT', '191', '5', '20');




INSERT INTO `register_course_check`.`subject`
(`id`,`subject_name`,`num_credits`,`faculty`)
VALUES
('CO1', 'AAA', '3', 'MT'),
('CO2', 'BBB', '4', 'HH'),
('CO3', 'CCC', '5', 'MT'),
('CO4', 'DDD', '3', 'CK'),
('CO5', 'New', '2', 'MT'),
('CO6', 'Test', '3', 'MT');






INSERT INTO `register_course_check`.`subject_condition`
(`id`,`subject_id`,`condition`)
VALUES
('1', 'CO1', '{\"data\": \"OR\", \"left\": {\"data\": \"CO2-1\"}, \"right\": {\"data\": \"AND\", \"left\": {\"data\": \"AND\", \"left\": {\"data\": \"CO5-1\"}, \"right\": {\"data\": \"CO6-2\"}}, \"right\": {\"data\": \"CO4-3\"}}}'),
('5', 'CO2', '{\"data\": \"CO3-1\"}'),
('6', 'CO3', '{\"data\": \"AND\", \"left\": {\"data\": \"CO4-1\"}, \"right\": {\"data\": \"CO5-1\"}}'),
('7', 'CO4', '{\"data\": \"CO2-3\"}');





