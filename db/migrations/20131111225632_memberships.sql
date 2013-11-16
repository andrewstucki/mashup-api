
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE memberships (
    github_account_id integer NOT NULL,
    user_id integer NOT NULL,
    is_admin boolean DEFAULT false,
    PRIMARY KEY (github_account_id, user_id)
);

CREATE INDEX index_memberships_on_github_account_id_and_user_id ON memberships USING btree (github_account_id, user_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP INDEX index_memberships_on_github_account_id_and_user_id;
DROP TABLE memberships;