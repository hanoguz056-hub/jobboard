CREATE TABLE IF NOT EXISTS "users" (

	"id" UUID DEFAULT gen_random_uuid() NOT NULL,

	"email" VARCHAR(255) NOT NULL UNIQUE,

	"password_hash" VARCHAR(255) NOT NULL,

	"role" VARCHAR(20),

	"is_banned" BOOLEAN DEFAULT false,

	"created_at" TIMESTAMPTZ DEFAULT NOW(),

	PRIMARY KEY("id")

);

CREATE TABLE IF NOT EXISTS "companies" (

	"id" UUID DEFAULT gen_random_uuid() NOT NULL,

	"owner_id" UUID NOT NULL,

	"name" VARCHAR(255) NOT NULL,

	"description" TEXT,

	"website" VARCHAR(255),

	"city" VARCHAR(255),

	"created_at" TIMESTAMPTZ DEFAULT NOW(),

	PRIMARY KEY("id")

);

CREATE TABLE IF NOT EXISTS "jobs" (

	"id" UUID DEFAULT gen_random_uuid() NOT NULL,

	"company_id" UUID,

	"title" VARCHAR(255),

	"description" TEXT,

	"salary_min" NUMERIC(10, 2),

	"salary_max" NUMERIC(10,2),

	"type" VARCHAR(20),

	"status" VARCHAR(20),

	"city" VARCHAR(100),

	"created_at" TIMESTAMPTZ DEFAULT NOW(),

	PRIMARY KEY("id")

);

CREATE TABLE IF NOT EXISTS "applications" (

	"id" UUID DEFAULT gen_random_uuid() NOT NULL,

	"job_id" UUID,

	"candidate_id" UUID,

	"resume_url" VARCHAR(500),

	"cover_letter" TEXT,

	"status" VARCHAR(20),

	"created_at" TIMESTAMPTZ DEFAULT NOW(),

	PRIMARY KEY("id")

);

ALTER TABLE "companies"

ADD FOREIGN KEY("owner_id") REFERENCES "users"("id")

ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "jobs"

ADD FOREIGN KEY("company_id") REFERENCES "companies"("id")

ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "applications"

ADD FOREIGN KEY("job_id") REFERENCES "jobs"("id")

ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "applications"

ADD FOREIGN KEY("candidate_id") REFERENCES "users"("id")

ON UPDATE NO ACTION ON DELETE NO ACTION;

ALTER TABLE "applications"

ADD CONSTRAINT unique_application

UNIQUE("job_id", "candidate_id")

