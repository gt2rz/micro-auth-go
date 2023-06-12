CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) NOT NULL UNIQUE COMMENT 'Email user',
        password VARCHAR(255) NOT NULL COMMENT 'Password user',
        firstname VARCHAR(255) NOT NULL COMMENT 'Firstname user',
        lastname VARCHAR(255) NOT NULL COMMENT 'Lastname user',
        phone VARCHAR(15) NOT NULL UNIQUE COMMENT 'Phone number in E.164 format',
        verified BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'True = Verified, False = Unverified',
        status BOOLEAN NOT NULL DEFAULT TRUE COMMENT 'True = Active, False = Inactive',
        password_reset_token_at TIMESTAMP DEFAULT NULL COMMENT 'Password reset token expiration date',
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );