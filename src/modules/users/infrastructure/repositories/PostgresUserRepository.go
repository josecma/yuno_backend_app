package repositories

import (
	"context"
	"yuno/src/modules/users/domain/entities"
	"yuno/src/modules/users/domain/repositories"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repositories.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *entities.User) error {

	dbUser := toDBModel(user)

	return r.db.WithContext(ctx).Save(dbUser).Error
}

func (r *PostgresUserRepository) RetrieveByID(ctx context.Context, id string) (*entities.User, error) {
	var dbUser UserModel

	// Buscar en DB
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&dbUser).Error

	if err != nil {
		return nil, err
	}

	// Convertir modelo DB a entidad de dominio
	return toDomainEntity(&dbUser), nil
}

// RetrieveByEmail implementa el método de la interfaz
func (r *PostgresUserRepository) RetrieveByEmail(ctx context.Context, email string) (*entities.User, error) {
	var dbUser UserModel

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&dbUser).Error

	if err != nil {
		return nil, err
	}

	return toDomainEntity(&dbUser), nil
}

// toDBModel convierte entidad dominio → modelo DB
func toDBModel(user *entities.User) *UserModel {
	return &UserModel{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
}

// toDomainEntity convierte modelo DB → entidad dominio
func toDomainEntity(dbUser *UserModel) *entities.User {
	return &entities.User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
