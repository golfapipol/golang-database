-- สร้าง database
CREATE SCHEMA `sckseal` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci ;

-- สร้าง table
CREATE TABLE `sckseal`.`sayhi` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`id`));

-- ใส่ข้อมูลลงในฐานข้อมูล
INSERT INTO `sckseal`.`sayhi` (`description`) VALUES ('hello world');