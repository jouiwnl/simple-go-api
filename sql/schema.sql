CREATE TABLE IF NOT EXISTS users (
    id varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255),
    cpf varchar(11),
    created_at timestamp(6),
    updated_at timestamp(6)
);
