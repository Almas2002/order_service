package model

import "7kzu-order-service/pkg/constans"

type Address struct {
	ID        uint32 `json:"id"`
	CityId    uint32 `json:"city_id"`
	Floor     uint32 `json:"floor"`
	Address   string `json:"address"`
	Apartment string `json:"apartment"`
	Entrance  string `json:"entrance"`
	Type      string `json:"type"`
}

func (a *Address) IsSender() bool {
	return constans.SENDER == a.Type
}
func (a *Address) IsReceiver() bool {
	return constans.RECEIVER == a.Type
}
