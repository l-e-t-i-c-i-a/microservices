package api

import (
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/ports"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// 1. Chama a porta do banco de dados para salvar
	err := a.db.Save(&order)
	// 2. Verifica se houve erro
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}
	// 3. Retorna o pedido salvo
	return order, nil
}