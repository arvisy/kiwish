CREATE TABLE "roles" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "role_id" int NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "address_id" int NOT NULL
);

CREATE TABLE "address" (
  "id" serial PRIMARY KEY,
  "address" varchar NOT NULL,
  "regency" varchar NOT NULL,
  "city" varchar NOT NULL
);


ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");
