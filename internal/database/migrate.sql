CREATE TYPE Status AS ENUM (
    'WAITING',
    'RUNNING',
    'SUCCEEDED',
    'FAILED'
    );

CREATE TABLE "Illustrator"
(
    "id"        SERIAL PRIMARY KEY,
    "name"      VARCHAR(255) UNIQUE NOT NULL,
    "pixivId"   VARCHAR(255) UNIQUE NOT NULL,
    "twitterId" VARCHAR(255) UNIQUE NOT NULL,
    "createdAt" TIMESTAMP DEFAULT (now()),
    "updatedAt" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "Illustration"
(
    "id"            SERIAL PRIMARY KEY,
    "title"         VARCHAR(255)        NOT NULL,
    "source"        VARCHAR(255) UNIQUE NOT NULL,
    "file"          VARCHAR(255)        NOT NULL,
    "createdAt"     TIMESTAMP DEFAULT (now()),
    "updatedAt"     TIMESTAMP DEFAULT (now()),
    "userId"        UUID                NOT NULL,
    "illustratorId" INT                 NOT NULL
);

CREATE TABLE "Tag"
(
    "id"             SERIAL PRIMARY KEY,
    "name"           VARCHAR(255) UNIQUE NOT NULL,
    "createdAt"      TIMESTAMP DEFAULT (now()),
    "updatedAt"      TIMESTAMP DEFAULT (now()),
    "illustrationId" INT                 NOT NULL
);

CREATE TABLE "Queue"
(
    "id"            SERIAL PRIMARY KEY,
    "source"        VARCHAR(255) UNIQUE NOT NULL,
    "status"        Status DEFAULT 'WAITING',
    "issuerId"      UUID NOT NULL,
    "createdAt"     TIMESTAMP DEFAULT (now()),
    "updatedAt"     TIMESTAMP DEFAULT (now()),
    "illustratorId" INT                 NOT NULL
);

ALTER TABLE "Illustration"
    ADD FOREIGN KEY ("illustratorId") REFERENCES "Illustrator" ("id");

ALTER TABLE "Tag"
    ADD FOREIGN KEY ("illustrationId") REFERENCES "Illustration" ("id");

ALTER TABLE "Queue"
    ADD FOREIGN KEY ("illustratorId") REFERENCES "Illustrator" ("id");
