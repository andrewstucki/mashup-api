
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
    
CREATE TABLE users (
    id integer NOT NULL PRIMARY KEY DEFAULT nextval('users_id_seq'),
    name character varying(255),
    login character varying(255),
    email character varying(255),
    encrypted_password character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

ALTER SEQUENCE users_id_seq OWNED BY users.id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP SEQUENCE users_id_seq;
DROP TABLE users;
