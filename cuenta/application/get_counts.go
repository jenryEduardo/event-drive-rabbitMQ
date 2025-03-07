package application

import "rabbitMQ/cuenta/domain"


type getCount struct {
	repo domain.Icuenta
}


func NewGetCount(repo domain.Icuenta)*getCount{
	return &getCount{repo: repo}
}


func(gc *getCount)Execute()([]domain.Cuenta,error){
	return gc.repo.GetAll()
}