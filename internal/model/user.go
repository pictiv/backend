package model

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type UserDTO struct {
	ID   uuid.UUID     `db:"id"`
	Name zeronull.Text `db:"name"`
	Role Role          `db:"role"`
}
