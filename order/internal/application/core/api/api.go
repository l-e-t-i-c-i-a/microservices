package api

import (
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/ports"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// Verifica se a quantidade total de itens excede 50
	if order.TotalQuantity() > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Total quantity of items cannot exceed 50")
	}

	// 1. Chama a porta do banco de dados para salvar (Status: Pending)
	err := a.db.Save(&order)
	// 2. Verifica se houve erro
	if err != nil {
		return domain.Order{}, err
	}
	// Tenta realizar o pagamento
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		// CASO DE ERRO: Atualiza status para Canceled
		order.Status = "Canceled"
		a.db.Update(order)
		// Retorna o erro para o cliente saber que falhou
		return domain.Order{}, paymentErr
	}

	// CASO DE SUCESSO: Atualiza status para Paid
	order.Status = "Paid"
	a.db.Update(order)

	// 3. Retorna o pedido
	return order, nil
}