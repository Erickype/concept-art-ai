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

// Get returns a user based on its email
//
//encore:api public method=GET path=/user/:email
func Get(ctx context.Context, email string) (*User, error) {
	u := &User{Email: email}
	err := sqldb.QueryRow(ctx, `
		SELECT user_id FROM "user"
		WHERE email = $1
	`, email).Scan(&u.Id)
	return u, err
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
