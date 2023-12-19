CREATE TABLE accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    `number` VARCHAR(8) UNIQUE NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) UNSIGNED DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `admin` BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    amount DECIMAL(10, 2) NOT NULL,
    transfered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    from_account_number VARCHAR(8) DEFAULT null,
    to_account_number VARCHAR(8) NOT NULL,
    FOREIGN KEY (from_account_number) REFERENCES accounts (number),
    FOREIGN KEY (to_account_number) REFERENCES accounts (number)
);
