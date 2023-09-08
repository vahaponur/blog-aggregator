-- name: CreatePost :one
INSERT INTO posts(id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;
-- name: GetPostsByUser :many
SELECT  * FROM posts
WHERE feed_id IN (SELECT feed_id FROM feed_follows where user_id = $1)
ORDER BY updated_at DESC ;

-- name: CheckPostExists :one
SELECT EXISTS (
    SELECT 1
    FROM posts
    WHERE url = $1
) AS post_exists;