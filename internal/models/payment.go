package models

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       uint   `json:"amount" validate:"gt=0"`
	PaymentDT    uint64 `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"gt=0"`
	GoodsTotal   int    `json:"goods_total" validate:"gt=0"`
	CustomFee    int    `json:"custom_fee" validate:"gt=0"`
}
