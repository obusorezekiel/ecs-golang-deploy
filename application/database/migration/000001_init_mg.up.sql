-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS id_seq;

-- Table Definition
CREATE TABLE "public"."students" (
    "ID" int4 NOT NULL DEFAULT nextval('id_seq'::regclass),
    "name" varchar,
    "description" varchar,
    PRIMARY KEY ("ID")
);

