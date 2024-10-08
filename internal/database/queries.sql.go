// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addToWaitlist = `-- name: AddToWaitlist :one
INSERT INTO waitlisted(id, twitter_id, twitter_username, referrer, email, wallet_address)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, twitter_id, twitter_username, referrer, email, wallet_address
`

type AddToWaitlistParams struct {
	ID              uuid.UUID
	TwitterID       string
	TwitterUsername string
	Referrer        sql.NullString
	Email           string
	WalletAddress   string
}

func (q *Queries) AddToWaitlist(ctx context.Context, arg AddToWaitlistParams) (Waitlisted, error) {
	row := q.db.QueryRowContext(ctx, addToWaitlist,
		arg.ID,
		arg.TwitterID,
		arg.TwitterUsername,
		arg.Referrer,
		arg.Email,
		arg.WalletAddress,
	)
	var i Waitlisted
	err := row.Scan(
		&i.ID,
		&i.TwitterID,
		&i.TwitterUsername,
		&i.Referrer,
		&i.Email,
		&i.WalletAddress,
	)
	return i, err
}

const getApplicant = `-- name: GetApplicant :one
SELECT id FROM waitlisted
WHERE id = $1
`

func (q *Queries) GetApplicant(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getApplicant, id)
	err := row.Scan(&id)
	return id, err
}

const getApplicantByTwitterID = `-- name: GetApplicantByTwitterID :one
SELECT twitter_id FROM waitlisted
WHERE twitter_id = $1
`

func (q *Queries) GetApplicantByTwitterID(ctx context.Context, twitterID string) (string, error) {
	row := q.db.QueryRowContext(ctx, getApplicantByTwitterID, twitterID)
	var twitter_id string
	err := row.Scan(&twitter_id)
	return twitter_id, err
}

const getLeaderboard = `-- name: GetLeaderboard :many
SELECT referrer, COUNT(referrer) AS num_referrals FROM waitlisted
WHERE referrer IS NOT NULL
GROUP BY referrer
ORDER BY num_referrals DESC
LIMIT $1
`

type GetLeaderboardRow struct {
	Referrer     sql.NullString
	NumReferrals int64
}

func (q *Queries) GetLeaderboard(ctx context.Context, limit int32) ([]GetLeaderboardRow, error) {
	rows, err := q.db.QueryContext(ctx, getLeaderboard, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLeaderboardRow
	for rows.Next() {
		var i GetLeaderboardRow
		if err := rows.Scan(&i.Referrer, &i.NumReferrals); err != nil {
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

const getUserReferralCount = `-- name: GetUserReferralCount :one
SELECT COUNT(referrer) FROM waitlisted WHERE referrer = $1
`

func (q *Queries) GetUserReferralCount(ctx context.Context, referrer sql.NullString) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUserReferralCount, referrer)
	var count int64
	err := row.Scan(&count)
	return count, err
}
