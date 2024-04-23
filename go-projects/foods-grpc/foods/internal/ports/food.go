package ports

type Food struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Quantity uint32 `json:"quantity"`
}
