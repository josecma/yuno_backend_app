package mappers

import (
	"yuno/src/common/database/models"
	"yuno/src/modules/users/domain/entities"
)

func FromDatabaseModelToDomainEntity(dbUser *models.User) *entities.User {

	return &entities.User{
		ID:       dbUser.ID,
		Email:    dbUser.Email,
		Username: dbUser.Username,
		Password: dbUser.Password,
	}

}

func FromDomainEntityToDatabaseModel(user *entities.User) *models.User {

	return &models.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}

}
