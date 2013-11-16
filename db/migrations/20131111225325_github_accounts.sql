
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE github_accounts (
    id integer NOT NULL PRIMARY KEY,
    login character varying(255),
    is_syncing boolean,
    synced_at timestamp without time zone,
    github_oauth_token character varying(255),
    gravatar_id character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
  
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE github_accounts;