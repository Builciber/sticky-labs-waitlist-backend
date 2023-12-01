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

-- +goose up
CREATE INDEX IF EXISTS twitter_id_idx ON waitlisted(twitter_id);

-- +goose down
DROP INDEX IF EXISTS twitter_id_idx;

-- +goose up
CREATE INDEX IF EXISTS referrer_idx ON waitlisted(referrer);

-- +goose down
DROP INDEX IF EXISTS referrer_idx;
