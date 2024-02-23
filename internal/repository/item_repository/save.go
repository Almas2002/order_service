package item_repository

import (
	"7kzu-order-service/internal/data/model"
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (r *repository) SaveItemTx(ctx context.Context, tx *sql.Tx, orderId string, item *model.Item) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "itemRepository.SaveItemTx")
	defer span.Finish()

	sql := `INSERT INTO items (order_id, image_url, title) VALUES ($1,$2,$3)`

	res, err := r.db.ExecContext(ctx, sql, orderId, item.ImageUrl, item.Title)
	if err != nil {
		return errors.Wrap(err, "SaveItemTx.ExecContext")
	}

	if _, err = res.RowsAffected(); err != nil {
		return errors.Wrap(err, "SaveItemTx.RowsAffected")
	}

	return err
}
