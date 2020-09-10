package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VitaminApp/database"
	"VitaminApp/graph/generated"
	"VitaminApp/graph/model"
	"context"
	"fmt"
	"time"
)

func (r *mutationResolver) CreateVitamin(ctx context.Context, input model.NewVitamin) (*model.Vitamin, error) {
	vitamin := &model.Vitamin{
		VitaminID:   fmt.Sprintf("T%d", 123),
		VitaminType: input.VitaminType,
		Benefits:    input.Benefits,
	}

	r.vitamins = append(r.vitamins, vitamin)
	return vitamin, nil
}

func (r *queryResolver) Vitamins(ctx context.Context) ([]*model.Vitamin, error) {
	query := fmt.Sprintf("SELECT VitaminID, VitaminType, Benefits FROM vitamin")
	context, cancel := context.WithTimeout(context.Background(), 8000*time.Millisecond)
	defer cancel()
	results, err := database.DbConn.QueryContext(context, query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	vitamins := make([]*model.Vitamin, 0)
	for results.Next() {
		var vitamin model.Vitamin
		results.Scan(&vitamin.VitaminID, &vitamin.VitaminType, &vitamin.Benefits)

		vitamins = append(vitamins, &vitamin)
	}
	return vitamins, nil
}

//Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
