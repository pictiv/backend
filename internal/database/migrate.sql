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
    "pixivId"   VARCHAR(255) UNIQUE,
    "twitterId" VARCHAR(255) UNIQUE,
    "createdAt" TIMESTAMPTZ DEFAULT (now()),
    "updatedAt" TIMESTAMPTZ DEFAULT (now())
);

CREATE TABLE "illustrations"
(
    "id"            SERIAL PRIMARY KEY,
    "title"         VARCHAR(255)        NOT NULL,
    "source"        VARCHAR(255) UNIQUE NOT NULL,
    "file"          VARCHAR(255)        NOT NULL,
    "createdAt"     TIMESTAMPTZ DEFAULT (now()),
    "updatedAt"     TIMESTAMPTZ DEFAULT (now()),
    "userId"        UUID                NOT NULL,
    "illustratorId" INT                 NOT NULL
);

CREATE TABLE "tags"
(
    "id"             SERIAL PRIMARY KEY,
    "name"           VARCHAR(255) UNIQUE NOT NULL,
    "createdAt"      TIMESTAMPTZ DEFAULT (now()),
    "updatedAt"      TIMESTAMPTZ DEFAULT (now()),
    "illustrationId" INT                 NOT NULL
);

CREATE TABLE "queue"
(
    "id"            SERIAL PRIMARY KEY,
    "source"        VARCHAR(255) UNIQUE NOT NULL,
    "status"        Status      DEFAULT 'WAITING',
    "issuerId"      UUID                NOT NULL,
    "createdAt"     TIMESTAMPTZ DEFAULT (now()),
    "updatedAt"     TIMESTAMPTZ DEFAULT (now()),
    "illustratorId" INT                 NOT NULL
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
    ADD FOREIGN KEY ("illustratorId") REFERENCES illustrators ("id");

ALTER TABLE tags
    ADD FOREIGN KEY ("illustrationId") REFERENCES illustrations ("id");

ALTER TABLE queue
    ADD FOREIGN KEY ("illustratorId") REFERENCES illustrators ("id");
