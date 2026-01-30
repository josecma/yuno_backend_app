package container

import (
	"yuno/src/common/config"
	"yuno/src/modules/users/domain/repositories"

	"gorm.io/gorm"
)

type Container struct {
	Config *config.Config
	db     *gorm.DB

	UserRepository repositories.UserRepository
}
