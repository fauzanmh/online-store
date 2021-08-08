BEGIN;

DROP TABLE IF EXISTS "products" CASCADE;

CREATE TABLE IF NOT EXISTS "products" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "price" numeric NOT NULL,
    "stock" int NOT NULL,
    "created_at" bigint NOT NULL,
    "updated_at" bigint
);

COMMIT;