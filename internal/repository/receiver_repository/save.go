package receiver_repository

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/tracing"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (r *receiverRepository) SaveReceiverTx(ctx context.Context, receiver *model.Receiver, orderId uuid.UUID, tx *sql.Tx) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "receiverRepository.SaveReceiverTx")
	defer span.Finish()

	addressId, err := r.adrRepo.SaveAddressTx(ctx, receiver.Address, tx)
	if err != nil {
		return err
	}

	fmt.Println(addressId)

	sql := `INSERT INTO receivers (order_id, person_name, address_id, phone) VALUES ($1,$2,$3,$4)`

	res, err := r.db.ExecContext(ctx, sql, orderId.String(), receiver.Name, addressId, receiver.Phone)
	if err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveReceiverTx.QueryxContext")
	}

	if _, err = res.RowsAffected(); err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveReceiverTx.RowsAffected")
	}
	logger.Infof("created sender with order id %s", orderId.String())

	return nil
}
