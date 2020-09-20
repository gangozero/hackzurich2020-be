package server

import (
	"context"
	"fmt"
	"log"

	"github.com/gangozero/hackzurich2020-be/generated/models"
	"github.com/jackc/pgx/v4"
)

func (s *Server) GetRecipeMatchList(ctx context.Context, userID string) (models.RecipeMatchList, error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("Can't acquire DB connection from pool: %s", err.Error())
	}
	defer conn.Release()

	query := `SELECT id, name, state, is_full, is_colleague, distance
				FROM public.matching 
				WHERE details->'users'->>$1 IS NOT NULL;`

	result := []*models.RecipeMatch{}

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.RecipeMatchList(result), nil
		}
		log.Printf("Error getting list of matches: %s", err.Error())
		return nil, fmt.Errorf("Error getting list of matches")
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var state models.RecipeMatchState
		var isFull, isColleague bool
		var distance int64

		err = rows.Scan(&id, &name, &state, &isFull, &isColleague, &distance)
		if err != nil {
			log.Printf("Error getting list of matches: %s", err.Error())
			return nil, fmt.Errorf("Error getting list of matches")
		}

		result = append(result, &models.RecipeMatch{
			ID:          &id,
			Name:        &name,
			State:       state,
			IsFull:      isFull,
			IsColleague: isColleague,
			Distance:    distance,
		})
	}

	return models.RecipeMatchList(result), nil
}

type details struct {
	Users       map[string]string          `json:"users,omitempty"`
	Image       string                     `json:"image,omitempty"`
	Steps       string                     `json:"steps,omitempty"`
	Ingridients []*models.RecipeIngredient `json:"ingridients,omitempty"`
}

func (s *Server) GetRecipeDetails(ctx context.Context, userID, matchID string) (*models.RecipeDetails, error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("Can't acquire DB connection from pool: %s", err.Error())
	}
	defer conn.Release()

	query := `SELECT id, name, state, is_full, is_colleague, distance, details
				FROM public.matching 
				WHERE id = $1;`

	var id, name string
	var state models.RecipeMatchState
	var isFull, isColleague bool
	var distance int64
	var details details

	err = conn.QueryRow(ctx, query, matchID).Scan(&id, &name, &state, &isFull, &isColleague, &distance, &details)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error getting details of matcing: %s", err.Error())
		return nil, fmt.Errorf("Error getting details of matcing")
	}

	partner := ""
	for k, v := range details.Users {
		if k != userID {
			partner = v
		}
	}

	return &models.RecipeDetails{
		ID:              &id,
		Name:            &name,
		ImageURL:        details.Image,
		Steps:           details.Steps,
		Ingredients:     details.Ingridients,
		PartnerName:     partner,
		PartnerDistance: distance,
	}, nil
}
