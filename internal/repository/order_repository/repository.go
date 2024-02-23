package order_repository

import (
	"7kzu-order-service/internal/repository"
	"github.com/jmoiron/sqlx"
)

type orderRepository struct {
	db           *sqlx.DB
	receiverRepo repository.ReceiverRepository
	senderRepo   repository.SenderRepository
	itemRepo     repository.ItemRepository
}

func New(db *sqlx.DB, receiverRepo repository.ReceiverRepository, senderRepo repository.SenderRepository,
	itemRepo repository.ItemRepository) *orderRepository {
	return &orderRepository{db, receiverRepo, senderRepo, itemRepo}
}
