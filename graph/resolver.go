package graph

import "VitaminApp/graph/model"

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	vitamins []*model.Vitamin
}
