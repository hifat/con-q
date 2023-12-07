-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "username" character varying(100) NULL,
  "password" character varying(200) NULL,
  "name" character varying(100) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX "users_username_key" ON "public"."users" ("username");
-- Create "auths" table
CREATE TABLE "public"."auths" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "token" text NULL,
  "agent" character varying(100) NULL,
  "client_ip" text NULL,
  "expires_at" timestamptz NULL,
  "user_id" uuid NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_auths_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "auths_token_key" to table: "auths"
CREATE UNIQUE INDEX "auths_token_key" ON "public"."auths" ("token");
