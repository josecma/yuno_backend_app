package repositories

import (
	"context"
	"yuno/src/common/database/models"
	"yuno/src/modules/users/domain/entities"
	"yuno/src/modules/users/domain/repositories"
	"yuno/src/modules/users/infrastructure/mappers"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repositories.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *entities.User) error {

	dbUser := mappers.FromDomainEntityToDatabaseModel(user)

	return r.db.WithContext(ctx).Save(dbUser).Error
}

func (r *PostgresUserRepository) RetrieveByID(ctx context.Context, id string) (*entities.User, error) {

	var dbUser models.User

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&dbUser).Error

	if err != nil {
		return nil, err
	}

	return mappers.FromDatabaseModelToDomainEntity(&dbUser), nil

}

func (r *PostgresUserRepository) RetrieveByEmail(ctx context.Context, email string) (*entities.User, error) {

	var dbUser models.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&dbUser).Error

	if err != nil {
		return nil, err
	}

	return mappers.FromDatabaseModelToDomainEntity(&dbUser), nil

}
