package dto

import (
	"7kzu-order-service/internal/data/model"
	"github.com/google/uuid"
)

type UpdateOrderDto struct {
	ID        uuid.UUID `json:"id"`
	CourierID *string   `json:"courier_id"`
	Status    string    `json:"status"`
}

func (o *UpdateOrderDto) ToModel() *model.Order {
	return &model.Order{
		ID:        o.ID,
		Status:    o.Status,
		CourierID: o.CourierID,
	}
}

type Order struct {
	ID     string  `json:"order_id"`
	TypeId uint32  `json:"type_id"`
	Volume float32 `json:"volume"`
	Lat    float32 `json:"lat"`
	Lot    float32 `json:"lot"`
	From   uint32  `json:"from"`
	To     uint32  `json:"to"`
}
