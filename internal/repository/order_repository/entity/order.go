package entity

import (
	"7kzu-order-service/internal/data/model"
	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID `db:"id"`
	CourierID   *string   `db:"courier_id"`
	MapUrl      string    `db:"map_url"`
	PutUrl      string    `db:"put_url"`
	Status      string    `db:"status"`
	CompanyId   uint32    `db:"company_id"`
	Price       uint64    `db:"price"`
	TransportId uint32    `db:"transport_id"`
	IsDraft     bool      `db:"is_draft"`
	Code        string    `db:"code"`
}

func (o *Order) ToModel() *model.Order {
	return &model.Order{
		ID:          o.ID,
		CourierID:   o.CourierID,
		MapUrl:      o.MapUrl,
		PutUrl:      o.PutUrl,
		Status:      o.Status,
		CompanyId:   o.CompanyId,
		Price:       o.Price,
		TransportId: o.TransportId,
		IsDraft:     o.IsDraft,
	}
}
