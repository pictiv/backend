CREATE TYPE Status AS ENUM (
    'WAITING',
    'RUNNING',
    'SUCCEEDED',
    'FAILED'
    );

CREATE TABLE "illustrators"
(
    "id"        SERIAL PRIMARY KEY,
    "name"      VARCHAR(255) UNIQUE NOT NULL,
    "pixiv_id"   VARCHAR(255) UNIQUE,
    "twitter_id" VARCHAR(255) UNIQUE,
    "created_at" TIMESTAMPTZ DEFAULT (now()),
    "updated_at" TIMESTAMPTZ DEFAULT (now())
);

CREATE TABLE "illustrations"
(
    "id"            SERIAL PRIMARY KEY,
    "title"         VARCHAR(255)        NOT NULL,
    "source"        VARCHAR(255) UNIQUE NOT NULL,
    "file"          VARCHAR(255)        NOT NULL,
    "created_at"     TIMESTAMPTZ DEFAULT (now()),
    "updated_at"     TIMESTAMPTZ DEFAULT (now()),
    "userId"        UUID                NOT NULL,
    "illustrator_id" INT                 NOT NULL
);

CREATE TABLE "tags"
(
    "id"             SERIAL PRIMARY KEY,
    "name"           VARCHAR(255) UNIQUE NOT NULL,
    "created_at"      TIMESTAMPTZ DEFAULT (now()),
    "updated_at"      TIMESTAMPTZ DEFAULT (now()),
    "illustration_id" INT                 NOT NULL
);

CREATE TABLE "queue"
(
    "id"            SERIAL PRIMARY KEY,
    "source"        VARCHAR(255) UNIQUE NOT NULL,
    "status"        Status      DEFAULT 'WAITING',
    "issuer_id"      UUID                NOT NULL,
    "created_at"     TIMESTAMPTZ DEFAULT (now()),
    "updated_at"     TIMESTAMPTZ DEFAULT (now()),
    "illustrator_id" INT                 NOT NULL
);

CREATE TYPE Role AS ENUM (
    'ADMIN',
    'MODERATOR',
    'USER'
    );

CREATE TABLE "users"
(
    "id"   UUID PRIMARY KEY,
    "name" VARCHAR(255) UNIQUE,
    "role" Role DEFAULT 'USER'
);


ALTER TABLE illustrations
    ADD FOREIGN KEY ("illustrator_id") REFERENCES illustrators ("id");

ALTER TABLE tags
    ADD FOREIGN KEY ("illustration_id") REFERENCES illustrations ("id");

ALTER TABLE queue
    ADD FOREIGN KEY ("illustrator_id") REFERENCES illustrators ("id");
