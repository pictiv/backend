package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type IllustratorRead struct {
	ID   int `query:"id"`
	Page int `param:"page"`
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
	ID        pgtype.Int4
	Name      pgtype.Text
	PixivID   pgtype.Text
	TwitterID pgtype.Text
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}
