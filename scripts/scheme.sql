USE gowallet;

CREATE TABLE IF NOT EXISTS `wallets` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(30) NOT NULL UNIQUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `transactions` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `wallet_id` INT NOT NULL,
  `ticker` VARCHAR(30) NOT NULL,
  `operation` VARCHAR(4) NOT NULL,
  `quantity` FLOAT NOT NULL,
  `price` FLOAT NOT NULL,
  `date` DATETIME,
  PRIMARY KEY (`id`),
  CONSTRAINT fk_transactions_wallets
  FOREIGN KEY (`wallet_id`) REFERENCES `wallets`(`id`)
) ENGINE=InnoDB;