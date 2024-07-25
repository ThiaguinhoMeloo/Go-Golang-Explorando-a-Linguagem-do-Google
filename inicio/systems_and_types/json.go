package systemsandtypes

type ProductJson struct {
	ID    int      `json:"id"`
	Name  string   `json: "nome"`
	Price float64  `json: "preco"`
	Tags  []string `json: "tags"`
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	systemsandtypes "inicio/systems_and_types"
// )

// func main() {
// 	//struct para Json
// 	p1 := systemsandtypes.ProductJson{
// 		ID:    1,
// 		Name:  "Notebook",
// 		Price: 1899.90,
// 		Tags: []string{
// 			"Promoção",
// 			"Eletrônico",
// 		},
// 	}

// 	p1Json, _ := json.Marshal(p1)
// 	fmt.Println(string(p1Json))

// 	fmt.Println("\n--------------------------------------------------------")
// 	//Json para struct
// 	var p2 systemsandtypes.ProductJson
// 	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
// 	json.Unmarshal([]byte(jsonString), &p2)
// 	fmt.Println(p2.Tags[1])
// }
