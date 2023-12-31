-- Create "reset_passwords" table
CREATE TABLE "public"."reset_passwords" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "code" character varying(20) NULL,
  "agent" character varying(100) NULL,
  "client_ip" character varying(30) NULL,
  "is_used" boolean NULL DEFAULT false,
  "expires_at" timestamptz NULL,
  "user_id" uuid NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_reset_passwords_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_reset_passwords_deleted_at" to table: "reset_passwords"
CREATE INDEX "idx_reset_passwords_deleted_at" ON "public"."reset_passwords" ("deleted_at");
