package server

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

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

func getUserFromBearer(principal interface{}) (string, error) {
	header, ok := principal.(string)
	if !ok {
		return "", fmt.Errorf("error casting interface{} to string")
	}
	data := strings.Split(header, " ")
	if len(data) != 2 {
		return "", fmt.Errorf("Wrong format of auth header")
	}

	return data[1], nil
}
