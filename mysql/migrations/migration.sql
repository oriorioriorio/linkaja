
-- create table
CREATE TABLE linkaja.customers (
	customer_number BIGINT NOT NULL auto_increment,
	name varchar(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
	PRIMARY KEY (`customer_number`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE linkaja.accounts (
	account_number BIGINT NOT NULL auto_increment,
	customer_number BIGINT NOT NULL,
	balance DECIMAL(18,2) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
	PRIMARY KEY (`account_number`),
	FOREIGN KEY (`customer_number`) REFERENCES customers(`customer_number`) ON UPDATE CASCADE ON DELETE CASCADE
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;


insert into customers(name) values ("agus"),("budi"),("tuti");

insert into accounts(customer_number, balance) values (1, 1000000.00),(2, 2000000.00),(3, 500000);
