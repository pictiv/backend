package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type IllustratorRead struct {
	ID   int `param:"id"`
	Page int `query:"page"`
}

type IllustratorSearch struct {
	ID        int    `query:"id"`
	Name      string `query:"name"`
	PixivID   string `query:"pixivId"`
	TwitterID string `query:"twitterId"`
}

type IllustratorDTO struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	PixivID   string    `db:"pixivId"`
	TwitterID string    `db:"twitterId"`
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

type IllustratorData struct {
	ID        pgtype.Int4        `db:"id"`
	Name      pgtype.Text        `db:"name"`
	PixivID   pgtype.Text        `db:"pixivId"`
	TwitterID pgtype.Text        `db:"twitterId"`
	CreatedAt pgtype.Timestamptz `db:"createdAt"`
	UpdatedAt pgtype.Timestamptz `db:"updatedAt"`
}
