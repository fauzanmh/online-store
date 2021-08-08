BEGIN;

ALTER TABLE "order_details" ADD CONSTRAINT "ID_PKEY" PRIMARY KEY ("id", "order_id");

COMMIT;