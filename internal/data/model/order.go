package model

import "github.com/google/uuid"

type Order struct {
	ID          uuid.UUID   `json:"id"`
	Receiver    *Receiver   `json:"receiver"`
	Sender      *Sender     `json:"sender"`
	CourierID   *string     `json:"courier_id"`
	MapUrl      string      `json:"map_url"`
	PutUrl      string      `json:"put_url"`
	Status      string      `json:"status"`
	CompanyId   uint32      `json:"company_id"`
	Price       uint64      `json:"price"`
	TransportId uint32      `json:"transport_id"`
	Items       []*Item     `json:"items"`
	Dimensions  *Dimensions `json:"dimensions"`
	IsDraft     bool        `json:"is_draft"`
	Code        string      `json:"code,omitempty"`
}

func (o *Order) IsFoundCourier() bool {
	return o.Status == "found courier"
}

func (o *Order) IsInShop() bool {
	return o.Status == "in shop"
}

func (o *Order) IsInHome() bool {
	return o.Status == "in home"
}

func (o *Order) IsSuccess() bool {
	return o.Status == "success"
}
