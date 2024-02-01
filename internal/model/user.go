package model

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type Role string

const (
	ADMIN     Role = "ADMIN"
	MODERATOR Role = "MODERATOR"
	USER      Role = "USER"
)

type UserDTO struct {
	ID   uuid.UUID
	Name zeronull.Text
	Role string
}
