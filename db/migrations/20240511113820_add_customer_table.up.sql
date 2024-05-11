CREATE TABLE customer
(
    id UUID NOT NULL PRIMARY KEY,
    phone_number VARCHAR(16) NOT NULL UNIQUE check ( char_length(phone_number) >= 10 AND char_length(phone_number) <= 16),
    name VARCHAR(50) NOT NULL check ( char_length(name) >= 5 AND char_length(name) <= 50),
    created_at timestamptz default now()
);
