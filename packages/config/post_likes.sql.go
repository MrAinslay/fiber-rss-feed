// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: post_likes.sql

package config

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPostLike = `-- name: CreatePostLike :one
INSERT INTO post_likes(id, created_at, updated_at, user_id, post_id)
VALUES($1, $2, $2, $3, $4)
RETURNING id, created_at, updated_at, user_id, post_id
`

type CreatePostLikeParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UserID    uuid.UUID
	PostID    uuid.UUID
}

func (q *Queries) CreatePostLike(ctx context.Context, arg CreatePostLikeParams) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, createPostLike,
		arg.ID,
		arg.CreatedAt,
		arg.UserID,
		arg.PostID,
	)
	var i PostLike
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const deletePostLike = `-- name: DeletePostLike :one
DELETE FROM post_likes WHERE id = $1
RETURNING id, created_at, updated_at, user_id, post_id
`

func (q *Queries) DeletePostLike(ctx context.Context, id uuid.UUID) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, deletePostLike, id)
	var i PostLike
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const getPostLikeById = `-- name: GetPostLikeById :one
SELECT id, created_at, updated_at, user_id, post_id FROM post_likes WHERE id = $1
`

func (q *Queries) GetPostLikeById(ctx context.Context, id uuid.UUID) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, getPostLikeById, id)
	var i PostLike
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.PostID,
	)
	return i, err
}

const getPostLikesByUser = `-- name: GetPostLikesByUser :many
SELECT id, created_at, updated_at, user_id, post_id FROM post_likes WHERE user_id = $1
`

func (q *Queries) GetPostLikesByUser(ctx context.Context, userID uuid.UUID) ([]PostLike, error) {
	rows, err := q.db.QueryContext(ctx, getPostLikesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostLike
	for rows.Next() {
		var i PostLike
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.PostID,
		); err != nil {
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
