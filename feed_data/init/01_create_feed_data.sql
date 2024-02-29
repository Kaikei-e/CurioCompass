CREATE SCHEMA IF NOT EXISTS feed_data;
CREATE TABLE IF NOT EXISTS feed_data.users (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);