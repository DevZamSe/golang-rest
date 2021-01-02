package entities

type Product struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	Status   bool    `json:"status"`
}
