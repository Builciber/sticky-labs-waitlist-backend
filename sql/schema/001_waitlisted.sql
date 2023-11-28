-- +goose Up
CREATE TABLE waitlisted (
    id UUID PRIMARY KEY,
    twitter_id VARCHAR(64) NOT NULL UNIQUE,
    twitter_username VARCHAR(64) NOT NULL UNIQUE,
    referrer VARCHAR(64),
    email VARCHAR(64) NOT NULL UNIQUE,
    wallet_address VARCHAR(64) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE waitlisted;