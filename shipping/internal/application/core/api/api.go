package api

import (
	"github.com/l-e-t-i-c-i-a/microservices/shipping/internal/application/core/domain"
)

type Application struct {
    // Shipping não precisa de banco de dados neste exemplo simples, 
    // mas se precisasse, injetaria db ports.DBPort aqui.
}

func NewApplication() *Application {
	return &Application{}
}

func (a Application) CreateShipping(shipping domain.Shipping) (domain.Shipping, error) {
	// 1. Calcula os dias
	shipping.CalculateDelivery()
    
    // 2. Aqui você salvaria no banco se fosse necessário.
    // Como o exercício não pediu explicitamente persistência no shipping,
    // vamos apenas retornar o cálculo.
	return shipping, nil
}