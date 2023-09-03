package models

type Item struct {
	ChrtID      uint64 `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required,min=13,max=13"`
	Price       uint64 `json:"price" validate:"gt=0"`
	Rid         string `json:"rid" validate:"required,min=21,max=36"`
	Name        string `json:"name" validate:"required"`
	Sale        uint64 `json:"sale" validate:"gt=0"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  uint64 `json:"total_price" validate:"gt=0"`
	NmId        uint64 `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      uint8  `json:"status" validate:"required,min=0, max=999"`
	OrderUid    string `json:"order_uid,omitempty"`
}
