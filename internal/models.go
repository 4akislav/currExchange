package internal

type Currency struct {
	Ccy      string  `json:"ccy"`
	Base_ccy string  `json:"base_ccy"`
	Buy      float64 `json:"buy,string"`
	Sale     float64 `json:"sale,string"`
}
