package data

type Pastry struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Ppu     float64   `json:"ppu"`
	Batters Batters   `json:"batters"`
	Topping []Topping `json:"topping"`
}

type Batter struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type Batters struct {
	Batter []Batter `json:"batter"`
}
type Topping struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
