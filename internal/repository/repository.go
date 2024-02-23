package repository

import (
	"7kzu-order-service/internal/data/model"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, model *model.Order) error
	UpdateOrder(ctx context.Context, id uuid.UUID, updateFn func(o *model.Order) *model.Order) error
}

type AddressRepository interface {
	SaveAddressTx(ctx context.Context, address *model.Address, tx *sql.Tx) (uint32, error)
}

type ReceiverRepository interface {
	SaveReceiverTx(ctx context.Context, receiver *model.Receiver, orderId uuid.UUID, tx *sql.Tx) error
}

type SenderRepository interface {
	SaveSenderTx(ctx context.Context, sender *model.Sender, orderId uuid.UUID, tx *sql.Tx) error
}
type ItemRepository interface {
	SaveItemTx(ctx context.Context, tx *sql.Tx, orderId string, item *model.Item) error
}
