package product

type CreateProductRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}
