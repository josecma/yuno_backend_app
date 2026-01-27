package container

import (
	"yuno/src/common/config"

	"gorm.io/gorm"
)

type Container struct {

	Config *config.Config
	db     *gorm.DB

	UserRepository         domain.UserRepository
}
