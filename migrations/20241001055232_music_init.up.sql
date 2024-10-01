CREATE TABLE songs (
    id VARCHAR(255) PRIMARY KEY,
    "group" VARCHAR(255),
    title VARCHAR(255),
    release_date DATE,
    text TEXT,
    patronymic VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
