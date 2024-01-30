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
  "address" varchar NOT NULL
);

CREATE TABLE "address" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "address" varchar NOT NULL,
  "regency" varchar NOT NULL,
  "city" varchar NOT NULL
);


ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "address" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
