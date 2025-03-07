package domain


type Icuenta interface {
	Save(Cuenta *Cuenta)error
	GetAll()([]Cuenta,error)
	Update(id int,cuenta *Cuenta)error
	Delete(id int)error
	Deposit(id int,mount float64)error
	Transfer(fromId int,toId int,mount float64)error
}