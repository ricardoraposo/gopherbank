CREATE TABLE accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `number` VARCHAR(8) UNIQUE NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) UNSIGNED DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `admin` BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE `users` (
    account_number VARCHAR(8) PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    FOREIGN KEY (account_number) REFERENCES accounts (number)
);

CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    from_account_number VARCHAR(8) DEFAULT null,
    to_account_number VARCHAR(8) DEFAULT null,
    FOREIGN KEY (from_account_number) REFERENCES accounts (number),
    FOREIGN KEY (to_account_number) REFERENCES accounts (number)
);

CREATE TABLE transaction_details (
    transaction_id INT PRIMARY KEY,
    amount DECIMAL(10, 2) NOT NULL,
    `type` VARCHAR(20) NOT NULL,
    transacted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions (id)
);
