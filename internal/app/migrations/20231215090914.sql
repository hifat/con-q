-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "fname", DROP COLUMN "lname", ADD COLUMN "name" character varying(100) NULL, ADD COLUMN "foo" character varying(100) NULL;
