package database

import (
	"context"
	"fmt"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"pictiv-api/internal/model"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
	Close()
	Health() bool
	Migrate() bool
	FindManyIllustrators(i model.IllustratorDTO, page int) ([]*model.IllustratorDTO, error)
	FindOneIllustrator(i model.IllustratorDTO) (model.IllustratorDTO, error)
	FindOneUser(i model.UserDTO) (model.UserDTO, error)
	CreateUser(i model.UserDTO) error
}

type service struct {
	db *pgxpool.Pool
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() Service {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	dbConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}
	db, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Close() {
	s.db.Close()
}

func (s *service) Health() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.Ping(ctx)
	if err != nil {
		return false
	}

	return true
}

func (s *service) Migrate() bool {
	_, err := s.db.Exec(context.Background(), `
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

		CREATE TABLE "users"
		(
			"id"      UUID PRIMARY KEY,
			"name"    VARCHAR(255) UNIQUE NOT NULL,
			"isAdmin" BOOL DEFAULT FALSE
		);

		ALTER TABLE illustrations
			ADD FOREIGN KEY ("illustratorId") REFERENCES illustrators ("id");
		
		ALTER TABLE tags
			ADD FOREIGN KEY ("illustrationId") REFERENCES illustrations ("id");
		
		ALTER TABLE queue
			ADD FOREIGN KEY ("illustratorId") REFERENCES illustrators ("id");
	`)
	if err != nil {
		return false
	}
	return true
}
