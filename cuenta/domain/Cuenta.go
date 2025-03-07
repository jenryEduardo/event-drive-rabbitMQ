package domain


type Cuenta struct {
	Titular string
	Saldo   float32 `json:"monto"`
	Moneda  string
	Creado_en string
}