package receiver_repository

import (
	"7kzu-order-service/internal/repository"
	"github.com/jmoiron/sqlx"
)

type receiverRepository struct {
	adrRepo repository.AddressRepository
	db      *sqlx.DB
}

func New(adrRepo repository.AddressRepository, db *sqlx.DB) *receiverRepository {
	return &receiverRepository{adrRepo, db}
}
