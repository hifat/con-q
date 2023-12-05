-- Modify "auths" table
ALTER TABLE "public"."auths" ADD COLUMN "expires_at" timestamptz NULL;
