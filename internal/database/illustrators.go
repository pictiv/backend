package database

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"pictiv-api/internal/model"
)

func (s *service) FindManyIllustrators(i model.IllustratorDTO, page int) ([]*model.IllustratorDTO, error) {
	j := model.IllustratorDTO{}
	var rows pgx.Rows
	var err error
	if i == j {
		// Empty
		rows, err = s.db.Query(context.Background(), `
		SELECT id, name, "twitterId", "pixivId", "createdAt", "updatedAt" FROM "Illustrator" LIMIT 10 OFFSET (10 * ($1 - 1));
	`, page)
		defer rows.Close()
	} else {
		// Has filters
		rows, err = s.db.Query(context.Background(), `
		SELECT id, name, "twitterId", "pixivId", "createdAt", "updatedAt" FROM "Illustrator" WHERE id=$1 OR name=$2 OR "pixivId"=$3 OR "twitterId"=$4 OR "createdAt"=$5 OR "updatedAt"=$6 LIMIT 10 OFFSET (10 * ($7 - 1));
	`, i.ID, i.Name, i.PixivID, i.TwitterID, i.CreatedAt, i.UpdatedAt, page)
		defer rows.Close()
	}
	if err != nil {
		return nil, err
	}
	var illustrator []*model.IllustratorDTO
	if err = pgxscan.ScanAll(&illustrator, rows); err != nil {
		return nil, err
	}
	return illustrator, nil
}

func (s *service) FindOneIllustrator(i model.IllustratorDTO) (model.IllustratorDTO, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id, name, "twitterId", "pixivId", "createdAt", "updatedAt" FROM "Illustrator" WHERE id=$1 OR name=$2 OR "pixivId"=$3 OR "twitterId"=$4 OR "createdAt"=$5 OR "updatedAt"=$6
		
	`, i.ID, i.Name, i.PixivID, i.TwitterID, i.CreatedAt, i.UpdatedAt)
	defer rows.Close()
	if err != nil {
		return model.IllustratorDTO{}, err
	}
	var illustrator model.IllustratorDTO
	if err = pgxscan.ScanOne(&illustrator, rows); err != nil {
		return model.IllustratorDTO{}, err
	}
	return illustrator, nil
}
