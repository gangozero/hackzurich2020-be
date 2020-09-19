package server

import "github.com/jackc/pgx/v4/pgxpool"

type Server struct {
	db *pgxpool.Pool
}

func NewServer(db *pgxpool.Pool) *Server {

	return &Server{
		db: db,
	}
}

// Validate function just return header that include userID
func (s *Server) Validate(header string) (interface{}, error) {
	return header, nil
}
