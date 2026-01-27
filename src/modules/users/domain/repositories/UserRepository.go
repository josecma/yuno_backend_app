package repositories

import (
	"context"
	"yuno/src/modules/users/domain/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user *entities.User) error
	RetrieveByID(ctx context.Context, id string) (*entities.User, error)
	RetrieveByEmail(ctx context.Context, email string) (*entities.User, error)
}
