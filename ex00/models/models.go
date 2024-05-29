package models

type Order struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type ThanksAndChange struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

type Error struct {
	Error string `json:"error"`
}
