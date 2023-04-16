package user

import (
	"context"
	"encore.dev/storage/sqldb"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type CreateRequest struct {
	Email string `json:"email"`
}

type CreateResponse struct {
	Rows   int64  `json:"rows"`
	Status string `json:"status"`
}

// Create creates a user receiving an email
//
//encore:api public method=POST path=/user
func Create(ctx context.Context, p *CreateRequest) (*CreateResponse, error) {
	rows, err := insert(ctx, p.Email)
	if err != nil {
		return nil, err
	}
	return &CreateResponse{Rows: rows, Status: "Created"}, nil
}

// insert inserts a URL into the database.
func insert(ctx context.Context, email string) (int64, error) {
	exec, err := sqldb.Exec(ctx, `
		INSERT INTO "user" (email)
		VALUES ($1)
	`, email)

	return exec.RowsAffected(), err
}
