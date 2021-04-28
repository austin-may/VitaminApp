package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"my-go-apps/VitaminApp/graph/generated"
	"my-go-apps/VitaminApp/graph/model"
	"my-go-apps/VitaminApp/vitamin"
)

func (r *mutationResolver) CreateVitamin(ctx context.Context, input model.NewVitamin) (*model.Vitamin, error) {
	err := vitamin.AddVitamin(input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) UpdateVitamin(ctx context.Context, input model.UpdatedVitamin) (*model.Vitamin, error) {
	err := vitamin.UpdateVitamin(input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) DeleteVitamin(ctx context.Context, vitaminID *int) (*model.Vitamin, error) {
	err := vitamin.DeleteVitamin(*vitaminID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *mutationResolver) SupplyVitamin(ctx context.Context, input model.SuppliedVitamin) (*int, error) {
	err := vitamin.SupplyVitamin(input)
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

func (r *queryResolver) VitaminByID(ctx context.Context, vitaminID *int) (*model.Vitamin, error) {
	vitamin, err := vitamin.GetVitaminById(*vitaminID)
	if err != nil {
		return nil, err
	}
	return vitamin, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
