package address_repository

import (
	"7kzu-order-service/internal/data/model"
	"7kzu-order-service/pkg/tracing"
	"context"
	"database/sql"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (r *repository) SaveAddressTx(ctx context.Context, address *model.Address, tx *sql.Tx) (uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "addressRepository.SaveAddressTx")
	defer span.Finish()

	var addressId uint32
	sqlCode := `INSERT INTO addresses (city_id, apartment, floor, entrance, address) VALUES ($1, $2, $3, $4, $5) RETURNING address_id`
	row, err := r.db.Query(sqlCode, address.CityId, address.Apartment, address.Floor, address.Entrance, address.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			tracing.TraceError(span, err)
			return 0, errors.New("No rows returned by the query")
		}
		tracing.TraceError(span, err)
		return 0, errors.Wrap(err, "SaveAddressTx.QueryRowxContext")
	}
	for row.Next() {
		if err = row.Scan(&addressId); err != nil {
			return 0, errors.Wrap(err, "SaveAddressTx")
		}
	}

	fmt.Println("addressId", addressId)

	return addressId, nil
}
