package application

import "rabbitMQ/cuenta/domain"

type deleteCount struct {
	repo domain.Icuenta
}


func NewDeleteCount(repo domain.Icuenta)*deleteCount{
	return &deleteCount{repo: repo}
}


func(CD *deleteCount)Execute(id int)error{
	return CD.repo.Delete(id)
}