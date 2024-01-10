-- Drop index "users_email_key" from table: "users"
DROP INDEX "public"."users_email_key";
-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "email" SET DEFAULT NULL::character varying;
