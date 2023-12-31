-- Modify "reset_passwords" table
ALTER TABLE "public"."reset_passwords" DROP COLUMN "is_used", ADD COLUMN "used_at" timestamptz NULL, ADD COLUMN "revoked_at" timestamptz NULL;
