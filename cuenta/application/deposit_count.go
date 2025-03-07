package application

import "rabbitMQ/cuenta/domain"

type Deposit struct {
	repo domain.Icuenta
}

func NewDeposit(repo domain.Icuenta)*Deposit{
	return &Deposit{repo: repo}
}


func(cd *Deposit)Execute(id int,mount float64)error{
	return cd.repo.Deposit(id,mount)
}