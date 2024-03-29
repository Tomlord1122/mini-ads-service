// Code generated by sqlc. DO NOT EDIT.
// source: ads.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createAds = `-- name: CreateAds :one
INSERT INTO ads (
   "title", "start_at", "end_at", "age", "gender", "country", "platform"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, title, start_at, end_at, age, gender, country, platform
`

type CreateAdsParams struct {
	Title    string         `json:"title"`
	StartAt  time.Time      `json:"start_at"`
	EndAt    time.Time      `json:"end_at"`
	Age      sql.NullInt32  `json:"age"`
	Gender   []GenderEnum   `json:"gender"`
	Country  []CountryEnum  `json:"country"`
	Platform []PlatformEnum `json:"platform"`
}

func (q *Queries) CreateAds(ctx context.Context, arg CreateAdsParams) (Ad, error) {
	row := q.db.QueryRowContext(ctx, createAds,
		arg.Title,
		arg.StartAt,
		arg.EndAt,
		arg.Age,
		pq.Array(arg.Gender),
		pq.Array(arg.Country),
		pq.Array(arg.Platform),
	)
	var i Ad
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.StartAt,
		&i.EndAt,
		&i.Age,
		pq.Array(&i.Gender),
		pq.Array(&i.Country),
		pq.Array(&i.Platform),
	)
	return i, err
}

const getActiveAds = `-- name: GetActiveAds :many


SELECT
    COUNT(*) 
FROM 
    ads
WHERE 
    start_at < NOW() AND end_at > NOW()
`

// $7是offset
func (q *Queries) GetActiveAds(ctx context.Context) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getActiveAds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			return nil, err
		}
		items = append(items, count)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAds = `-- name: ListAds :many
SELECT
    "title", "end_at"
FROM
    ads
WHERE
    start_at < NOW() AND end_at > NOW()
    AND ($1::int IS NULL OR age >= $1) -- $1是年齡下限
    AND ($2::int IS NULL OR age <= $2) -- $2是年齡上限
    AND ($3::country_enum[] IS NULL OR country && $3) -- $3是國家條件
    AND ($4::platform_enum[] IS NULL OR platform && $4) -- $4是平台條件
    AND ($5::gender_enum[] IS NULL OR gender && $5) -- $5是性別條件
ORDER BY
    end_at ASC
LIMIT LEAST(GREATEST(COALESCE(NULLIF($6, 0), 5), 1), 100) -- $6是limit
OFFSET $7
`

type ListAdsParams struct {
	Column1 int32          `json:"column_1"`
	Column2 int32          `json:"column_2"`
	Column3 []CountryEnum  `json:"column_3"`
	Column4 []PlatformEnum `json:"column_4"`
	Column5 []GenderEnum   `json:"column_5"`
	Column6 interface{}    `json:"column_6"`
	Offset  int32          `json:"offset"`
}

type ListAdsRow struct {
	Title string    `json:"title"`
	EndAt time.Time `json:"end_at"`
}

func (q *Queries) ListAds(ctx context.Context, arg ListAdsParams) ([]ListAdsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAds,
		arg.Column1,
		arg.Column2,
		pq.Array(arg.Column3),
		pq.Array(arg.Column4),
		pq.Array(arg.Column5),
		arg.Column6,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAdsRow
	for rows.Next() {
		var i ListAdsRow
		if err := rows.Scan(&i.Title, &i.EndAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
