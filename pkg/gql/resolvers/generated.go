package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/ameernormie/go-api-template/pkg/gql"
	"github.com/ameernormie/go-api-template/pkg/gql/models"
)

type Resolver struct{}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	records := []*models.User{
		&models.User{
			ID:     "ec17af15-e354-440c-a09f-69715fc8b595",
			Email:  "test@mail.com",
			UserID: "user1",
		},
	}
	return records, nil

}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
