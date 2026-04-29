-- +migrate Up
CREATE TABLE IF NOT EXISTS "monogo"."admins" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
    "name" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz NULL
);

-- +migrate Down
DROP TABLE IF EXISTS "monogo"."admins";