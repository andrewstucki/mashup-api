
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE repositories (
    id integer NOT NULL PRIMARY KEY,
    name character varying(255),
    url character varying(255),
    owner_name character varying(255),
    active boolean DEFAULT false,
    description text,
    default_branch character varying(255),
    vimeo_key character varying(255),
    flickr_key character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE INDEX index_repositories_on_owner_name_and_name ON repositories USING btree (owner_name, name);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP INDEX index_repositories_on_owner_name_and_name;
DROP TABLE repositories;