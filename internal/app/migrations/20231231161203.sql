-- Modify "auths" table
ALTER TABLE "public"."auths" ALTER COLUMN "client_ip" TYPE character varying(30);
-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "email" character varying(100) NULL;
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "public"."users" ("email");
