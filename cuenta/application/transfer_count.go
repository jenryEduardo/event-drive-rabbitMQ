package application

import "rabbitMQ/cuenta/domain"

type Transfer struct {
	repo domain.Icuenta
}

func NewTransfer(repo domain.Icuenta)*Transfer{
	return &Transfer{repo: repo}
}

func(ct *Transfer)Execute(fromId int,toId int,mount float64)error{
	return ct.repo.Transfer(fromId,toId,mount)
}