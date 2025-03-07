package application

import "rabbitMQ/cuenta/domain"

type UpdateCount struct {
	repo domain.Icuenta
}


func NewUpdate(repo domain.Icuenta)*UpdateCount{
	return &UpdateCount{repo: repo}
}


func (cu *UpdateCount)Execute(id int,count domain.Cuenta)error{
	return cu.repo.Update(id,&count)
}