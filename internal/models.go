package internal

type Currency struct {
	Ccy      string  `json:"ccy"`
	Base_ccy string  `json:"base_ccy"`
	Buy      float32 `json:"buy"`
	Sale     float32 `json:"sale"`
}
