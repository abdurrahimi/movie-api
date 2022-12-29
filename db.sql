CREATE TABLE IF NOT EXISTS "public"."movie" (
  "id" SERIAL PRIMARY KEY,
  "title" text,
  "description" text,
  "rating" float,
  "image" text,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
)
;