BEGIN;

ALTER TABLE "order_details" DROP CONSTRAINT "ID_PKEY";

COMMIT;