package sender_repository

import (
	"7kzu-order-service/internal/repository"
	"github.com/jmoiron/sqlx"
)

type senderRepository struct {
	adrRepo repository.AddressRepository
	db      *sqlx.DB
}

func New(adrRepo repository.AddressRepository, db *sqlx.DB) *senderRepository {
	return &senderRepository{adrRepo, db}
}
