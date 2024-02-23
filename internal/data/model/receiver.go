package model

type Receiver struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address *Address
}
