package database

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"pictiv-api/internal/model"
)

func (s *service) FindOneUser(i model.UserDTO) (model.UserDTO, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id, name, role from users where id=$1
	`, i.ID)
	defer rows.Close()
	if err != nil {
		return model.UserDTO{}, err
	}

	var user model.UserDTO
	if err = pgxscan.ScanOne(&user, rows); err != nil {
		return model.UserDTO{}, err
	}
	if user.Name == "" {
		return model.UserDTO{
			ID:   i.ID,
			Name: zeronull.Text(fmt.Sprintf("%v", i.ID)),
			Role: i.Role,
		}, nil
	} else {
		return user, nil
	}
}

func (s *service) CreateUser(i model.UserDTO) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO users (id) VALUES ($1)
	`, i.ID)
	if err != nil {
		return err
	}
	return nil
}
