package order_repository

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/tracing"
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (r *orderRepository) SaveOrder(ctx context.Context, model *model.Order) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepository.SaveOrder")
	defer span.Finish()

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	if err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveOrder.BeginTxx")
	}

	sqlCode := "INSERT INTO orders (order_id, company_id,put_url, price, transport_id) VALUES ($1,$2,$3,$4,$5)"

	res, err := r.db.ExecContext(ctx, sqlCode, model.ID.String(), model.CompanyId, model.PutUrl, model.Price, model.TransportId)
	if err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveOrder.ExecContext")
	}

	if _, err = res.RowsAffected(); err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "SaveOrder.RowsAffected")
	}

	if err = r.receiverRepo.SaveReceiverTx(ctx, model.Receiver, model.ID, tx); err != nil {
		return err
	}

	if err = r.senderRepo.SaveSenderTx(ctx, model.Sender, model.ID, tx); err != nil {

	}

	//for i := 0; i < len(model.Items); i++ {
	//
	//	if err = r.itemRepo.SaveItemTx(ctx, tx, model.ID.String(), model.Items[i]); err != nil {
	//		return err
	//	}
	//
	//}

	logger.Infof("order created with id %s", model.ID.String())

	return nil
}
