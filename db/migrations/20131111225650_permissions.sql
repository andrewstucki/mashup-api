
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE permissions (
    user_id integer,
    repository_id integer,
    admin boolean DEFAULT false,
    PRIMARY KEY (repository_id, user_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE permissions;