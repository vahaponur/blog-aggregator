-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id,user_id,feed_id,created_at,updated_at)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;
-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = $1;
-- name: GetAllFeedsOfUser :many
SELECT * FROM feed_follows WHERE user_id=$1;
-- name: FeedFollowExist :one
SELECT EXISTS (
    SELECT 1
    FROM feed_follows
    WHERE user_id = $1 AND id = $2
) AS follow_exists;