-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY DEFAULT UUID(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    age INT CHECK (age > 0 AND age < 130),
    phone VARCHAR(20) UNIQUE,
    status ENUM('active', 'inactive', 'banned') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    -- Indexes
    INDEX idx_users_email (email),
    INDEX idx_users_status (status),
    INDEX idx_users_phone (phone)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
