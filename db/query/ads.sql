-- name: CreateAds :one
INSERT INTO ads (
   "title", "start_at", "end_at", "age", "gender", "country", "platform"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;



-- name: ListAds :many
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
OFFSET $7; -- $7是offset


-- name: GetActiveAds :many
SELECT
    COUNT(*) 
FROM 
    ads
WHERE 
    start_at < NOW() AND end_at > NOW();