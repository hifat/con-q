-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "name", ADD COLUMN "fname" character varying(100) NULL, ADD COLUMN "lname" character varying(100) NULL;
