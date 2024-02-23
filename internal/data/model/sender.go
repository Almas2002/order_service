package model

type Sender struct {
	Name        string   `json:"name"`
	Phone       string   `json:"phone"`
	CompanyName string   `json:"company_name"`
	Address     *Address `json:"address"`
}
