package models

type AutoWithdraw struct {
	PartnerID int  `json:"partner_id"`
	Enabled   bool `json:"enabled"`
}
