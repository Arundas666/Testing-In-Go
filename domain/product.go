package domain

type Product struct {
	ProductName string  `json:"productname"`
	Price       float64 `json:"price"`
	Catergory   string  `json:"category"`
}
 