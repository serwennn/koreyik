package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/serwennn/koreyik/internal/storage/pq"
	"time"
)

type Anime struct {
	ID           int
	ThumbnailURL string
	Description  string
	Rating       string

	TitleKk string
	TitleJp string
	TitleEn string

	Status         string
	StartedAiring  time.Time
	FinishedAiring time.Time

	Genres []string
	Themes []string

	Seasons  int
	Episodes int
	Duration int

	Studios   []string
	Producers []string

	//Related []MediaEntry
}

func CreateAnime(storage *pq.Storage, ctx context.Context, a Anime) error {
	_, err := storage.DB.Exec(
		ctx, "INSERT INTO animes VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)",
		a.ID, a.ThumbnailURL, a.Description, a.Rating,
		a.TitleKk, a.TitleJp, a.TitleEn,
		a.Status, a.StartedAiring, a.FinishedAiring,
		a.Genres, a.Themes,
		a.Seasons, a.Episodes,
		a.Duration, a.Studios, a.Producers,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetAnime(storage *pq.Storage, ctx context.Context, id int) (Anime, error) {
	row, _ := storage.DB.Query(ctx, "SELECT * FROM animes WHERE id = $1", id)

	anime, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Anime])
	if err != nil {
		return Anime{}, err
	}
	return anime, nil
}
