package repository

import (
	"behealth-api/infrastructure"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db infrastructure.Database,
	logger infrastructure.Logger) UserRepository {
	return UserRepository{
		Database: db,
		logger:   logger,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Zap.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
