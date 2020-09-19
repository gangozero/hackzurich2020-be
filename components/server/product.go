package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gangozero/hackzurich2020-be/generated/models"
	"github.com/jackc/pgx/v4"
)

func (s *Server) GetProductList(ctx context.Context, userID string) (models.ProductList, error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("Can't acquire DB connection from pool: %s", err.Error())
	}
	defer conn.Release()

	query := `SELECT productID, name
				FROM public.products 
				WHERE userID = $1;`

	result := []*models.Product{}

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.ProductList(result), nil
		}
		log.Printf("Error getting list of products: %s", err.Error())
		return nil, fmt.Errorf("Error getting list of products")
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			log.Printf("Error getting list of products: %s", err.Error())
			return nil, fmt.Errorf("Error getting list of products")
		}

		result = append(result, &models.Product{
			ID:   &id,
			Name: &name,
		})
	}

	return models.ProductList(result), nil
}

func (s *Server) DeleteProduct(ctx context.Context, userID string, productID string) error {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("Can't acquire DB connection from pool: %s", err.Error())
	}
	defer conn.Release()

	query := `DELETE
				FROM public.products 
				WHERE userID = $1 AND productID = $2;`

	_, err = conn.Exec(ctx, query, userID, productID)
	if err != nil {
		log.Printf("Error deleting product from list: %s", err.Error())
		return fmt.Errorf("Error deleting product from list")
	}

	return nil
}
