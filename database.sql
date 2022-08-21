CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "gmail" varchar,
  "phone_number" varchar,
  "name" varchar,
  "password" varchar,
  "dob" datetime,
  "type" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "tests" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial,
  "name" varchar,
  "subject_id" integer,
  "duration" duration,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "questions" (
  "id" bigserial PRIMARY KEY,
  "test_id" bigserial,
  "content" varchar,
  "images" varchar,
  "type" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "options" (
  "question_id" bigserial PRIMARY KEY,
  "text" varchar,
  "image" varchar,
  "is_answer" boolean,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "subjects" (
  "id" integer PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "student_test" (
  "id" bigserial,
  "student_id" bigserial,
  "test_id" bigserial,
  "start" timestamptz,
  "end" timestamptz,
  "created_at" timestamptz DEFAULT (now()),
  PRIMARY KEY ("student_id", "test_id")
);

CREATE TABLE "student_answer" (
  "student_test_id" bigserial,
  "question_id" bigserial,
  "option_id" bigserial,
  PRIMARY KEY ("student_test_id", "question_id")
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("phone_number");

CREATE INDEX ON "tests" ("name");

CREATE INDEX ON "tests" ("id");

CREATE INDEX ON "questions" ("id");

CREATE INDEX ON "questions" ("test_id");

CREATE INDEX ON "student_test" ("id");

CREATE INDEX ON "student_test" ("student_id");

CREATE INDEX ON "student_test" ("test_id");

COMMENT ON COLUMN "tests"."user_id" IS 'user_id create this test';

ALTER TABLE "tests" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tests" ADD FOREIGN KEY ("subject_id") REFERENCES "subjects" ("id");

ALTER TABLE "questions" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");

ALTER TABLE "options" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "student_test" ADD FOREIGN KEY ("student_id") REFERENCES "users" ("id");

ALTER TABLE "student_test" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");

ALTER TABLE "student_answer" ADD FOREIGN KEY ("student_test_id") REFERENCES "student_test" ("id");
