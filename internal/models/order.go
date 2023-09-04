package models

type Order struct {
	OrderUid          string   `json:"order_uid" validate:"required,min=19,max=36"`
	TrackNumber       string   `json:"track_number" validate:"required,min=13,max=14"`
	Entry             string   `json:"entry" validate:"required,min=4,max=4"`
	Delivery          Delivery `json:"delivery" validate:"required"`
	Payment           Payment  `json:"payment" validate:"required"`
	Items             []Item   `json:"items" validate:"required"`
	Locale            string   `json:"locale" validate:"oneof=ru en"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id" validate:"required,min=4,max=4"`
	DeliveryService   string   `json:"delivery_service" validate:"required,min=5,max=5"`
	ShardKey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id" validate:"gte=0,lte=100"`
	DateCreated       string   `json:"date_created" db:"date_created" format:"2021-11-26T06:22:19Z" validate:"required"`
	OofShard          string   `json:"oof_shard" validate:"required,max=2"`
}
