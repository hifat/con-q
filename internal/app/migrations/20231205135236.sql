-- Modify "auths" table
ALTER TABLE "public"."auths" DROP COLUMN "token", ADD COLUMN "refresh_token" text NULL;
