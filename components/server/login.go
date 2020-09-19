package server

import (
	"context"
	"fmt"

	"github.com/gangozero/hackzurich2020-be/generated/models"
	"github.com/jackc/pgx/v4"
)

func (s *Server) Login(ctx context.Context, email string) (*models.LoginResponse, error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("Can't acquire DB connection from pool: %s", err.Error())
	}
	defer conn.Release()

	query := `SELECT id
				FROM public.users 
				WHERE email = $1;`

	var id string
	err = conn.QueryRow(ctx, query, email).Scan(&id)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Can't get user ID: %s", err.Error())
	}

	return &models.LoginResponse{
		Token: id,
	}, nil
}
