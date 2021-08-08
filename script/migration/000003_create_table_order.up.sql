BEGIN;

DROP TABLE IF EXISTS "orders" CASCADE;

CREATE TABLE IF NOT EXISTS "orders" (
    "id" bigserial PRIMARY KEY,
    "user_id" UUID NOT NULL,
    "status" varchar(1) NOT NULL DEFAULT 0,
    "created_at" bigint NOT NULL,
    "updated_at" bigint
);

COMMIT;