package sender_repository

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/tracing"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (r *senderRepository) SaveSenderTx(ctx context.Context, sender *model.Sender, orderId uuid.UUID, tx *sql.Tx) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "senderRepository.SaveSenderTx")
	defer span.Finish()

	addressId, err := r.adrRepo.SaveAddressTx(ctx, sender.Address, tx)
	if err != nil {
		return err
	}

	sql := `INSERT INTO senders (order_id, person_name, address_id, phone, company_name) VALUES ($1,$2,$3,$4,$5)`

	res, err := r.db.ExecContext(ctx, sql, orderId, sender.Name, addressId, sender.Phone, sender.CompanyName)
	if err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveSenderTx.QueryxContext")
	}

	if _, err = res.RowsAffected(); err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveSenderTx.RowsAffected")
	}
	logger.Infof("created sender with order id %s", orderId.String())

	return nil
}
