CREATE TABLE "sellers" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "name" varchar NOT NULL,
  "address_id" int NOT NULL,
  "last_active" date NOT NULL
);

CREATE TABLE "address" (
  "id" serial PRIMARY KEY,
  "user_id" int NOT NULL,
  "address" varchar NOT NULL,
  "regency" varchar NOT NULL,
  "city" varchar NOT NULL
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "seller_id" int NOT NULL,
  "name" varchar NOT NULL,
  "price" decimal NOT NULL,
  "stock" int NOT NULL,
  "category_id" int NOT NULL,
  "discount" int NOT NULL DEFAULT 0
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

ALTER TABLE "sellers" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("seller_id") REFERENCES "sellers" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");