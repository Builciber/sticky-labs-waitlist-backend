-- name: AddToWaitlist :one
INSERT INTO waitlisted(id, twitter_id, twitter_username, referrer, email, wallet_address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetLeaderboard :many
SELECT referrer, COUNT(referrer) AS num_referrals FROM waitlisted
WHERE referrer IS NOT NULL
GROUP BY referrer
ORDER BY num_referrals DESC
LIMIT $1;

-- name: GetApplicant :one
SELECT id FROM waitlisted
WHERE id = $1;

-- name: GetApplicantByTwitterID :one
SELECT twitter_id FROM waitlisted
WHERE twitter_id = $1;

-- name: GetUserReferralCount :one
SELECT COUNT(referrer) FROM waitlisted WHERE referrer = $1;

-- name: CreateTwitterIDIndex :exec
CREATE INDEX twitter_id_idx ON waitlisted(twitter_id);

-- name: CreateReferrerIndex :exec
CREATE INDEX referrer_idx ON waitlisted(referrer);