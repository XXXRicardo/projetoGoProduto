package models

type Mercado struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	Nome       string  `json: "nome"`
	Descricao  string  `json:"descricao"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}
