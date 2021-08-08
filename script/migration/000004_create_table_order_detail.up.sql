BEGIN;

DROP TABLE IF EXISTS "order_details" CASCADE;

CREATE TABLE IF NOT EXISTS "order_details" (
    "id" bigint,
    "order_id" bigint,
    "product_id" bigint NOT NULL,
    "price" numeric NOT NULL,
    "qty" int NOT NULL,
    "total_price" numeric NOT NULL,
    "created_at" bigint NOT NULL,
    "updated_at" bigint
);

COMMIT;