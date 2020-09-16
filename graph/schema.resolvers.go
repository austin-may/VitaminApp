package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"VitaminApp/graph/generated"
	"VitaminApp/graph/model"
	"VitaminApp/vitamin"
	"context"
)

func (r *mutationResolver) CreateVitamin(ctx context.Context, input model.NewVitamin) (*model.Vitamin, error) {
	err := vitamin.AddVitamin(input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *queryResolver) Vitamins(ctx context.Context) ([]*model.Vitamin, error) {
	vitamins, err := vitamin.GetVitaminList()
	if err != nil {
		return nil, err
	}
	return vitamins, nil
}

//Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
