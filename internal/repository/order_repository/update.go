package order_repository

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/internal/repository/order_repository/entity"
	"7kzu-order-service/pkg/custom_errors"
	"7kzu-order-service/pkg/logger"
	"7kzu-order-service/pkg/tracing"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"time"
)

func (r *orderRepository) UpdateOrder(ctx context.Context, id uuid.UUID, updateFn func(o *model.Order) *model.Order) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepository.UpdateOrder")
	defer span.Finish()

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	defer func(err error) error {
		if err != nil {
			if err = tx.Rollback(); err != nil {
				return errors.Wrap(err, "UpdateOrder.Rollback")
			}
		}
		if err = tx.Commit(); err != nil {
			return errors.Wrap(err, "UpdateOrder.Commit")
		}
		return nil
	}(err)
	if err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "UpdateOrder.BeginTx")
	}

	oldOrder, err := r.findOneTx(ctx, id, tx)
	if err != nil {
		return err
	}
	newOrder := updateFn(oldOrder)

	sqlCode := `UPDATE orders SET courier_id = $1 , status = $2 , map_url= $3 , price = $4 , updated_at=$5 , company_id =$6, code =$7 WHERE order_id = $8`
	exec, err := tx.Exec(sqlCode, newOrder.CourierID, newOrder.Status, newOrder.MapUrl, newOrder.Price, time.Now(), newOrder.CompanyId, newOrder.Code, newOrder.ID.String())
	if err != nil {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "UpdateOrder.ExecContext")
	}

	if count, err := exec.RowsAffected(); err != nil && count == 0 {
		tracing.TraceError(span, err)
		return errors.Wrap(err, "UpdateOrder.RowsAffected")
	}

	logger.Infof("Updated order with id: %s", newOrder.ID.String())

	return nil
}

func (r *orderRepository) findOneTx(ctx context.Context, id uuid.UUID, tx *sql.Tx) (*model.Order, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "orderRepository.findOneTx")
	defer span.Finish()

	sqlCode := `SELECT order_id,company_id,courier_id,status,map_url,put_url,price,transport_id,code FROM orders WHERE order_id=$1`

	order := entity.Order{}

	err := tx.QueryRowContext(ctx, sqlCode, id.String()).Scan(&order.ID, &order.CompanyId, &order.CourierID, &order.Status, &order.MapUrl, &order.PutUrl, &order.Price, &order.TransportId, &order.Code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tracing.TraceError(span, err)
			return nil, errors.Wrap(custom_errors.NotFoundError, "findOneTx.QueryContext")
		}
		tracing.TraceError(span, err)
		return nil, errors.Wrap(err, "findOneTx.QueryContext")
	}

	logger.Infof("found order with id: %s", id.String())

	return order.ToModel(), nil

}
