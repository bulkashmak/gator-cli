-- name: GetFeedFollowsForUser :many
SELECT 
    ff.*,
    f.name AS feed_name,
    u.name AS user_name
FROM feed_followers ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;

-- name: CreateFeedFollower :one
WITH inserted_feed_follower AS (
  INSERT INTO feed_followers (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT inserted_feed_follower.*, feeds.name AS feed_name, users.name AS user_name
FROM inserted_feed_follower
JOIN feeds ON inserted_feed_follower.feed_id = feeds.id
JOIN users ON inserted_feed_follower.user_id = users.id;

