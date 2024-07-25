package systemsandtypes

type ProductJson struct {
	ID    int      `json:"id"`
	Name  string   `json: "nome"`
	Price float64  `json: "preco"`
	Tags  []string `json: "tags"`
}
