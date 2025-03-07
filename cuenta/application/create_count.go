package application

import "rabbitMQ/cuenta/domain"

type CreateCount struct {
	repo domain.Icuenta
}

func NewCreateCount(repo domain.Icuenta)*CreateCount{
	return &CreateCount{repo: repo}
}


func(cc *CreateCount)Execute(count domain.Cuenta)error{
	return cc.repo.Save(&count)
}
