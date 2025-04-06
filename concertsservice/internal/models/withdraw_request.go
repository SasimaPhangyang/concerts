package models

type WithdrawRequest struct {
	PartnerID int     `json:"partner_id"`
	Amount    float64 `json:"amount"`
}
