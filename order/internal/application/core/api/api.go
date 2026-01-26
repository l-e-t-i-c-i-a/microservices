package api

import (
	"log"

	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/ports"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db: db,
		payment: payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// Verifica se a quantidade total de itens excede 50
	if order.TotalQuantity() > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Total quantity of items cannot exceed 50")
	}

	// Validação de Estoque (NOVO)
    if err := a.db.CheckProductsExist(order.OrderItems); err != nil {
        return domain.Order{}, status.Errorf(codes.NotFound, err.Error())
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

	// Pagamento deu certo! Agora chamamos o Shipping (Requisito 1.1)
	// "apenas realize a requisição para o microsserviço Shipping, caso o pagamento ocorra com sucesso"
	shippingErr := a.shipping.ShipOrder(order)
	if shippingErr != nil {
		// Se der erro no envio, logamos o erro, mas o pedido já foi pago.
		// Decisão de projeto: Mantemos como "Paid" ou mudamos para algo como "ShippingFailed".
		// Vamos logar e retornar o erro para o cliente saber.
		log.Printf("❌ Erro ao solicitar envio: %v", shippingErr)
		order.Status = "Paid"
	} else {
		// Se o envio for solicitado com sucesso, atualizamos o status para "Shipped"
		order.Status = "Shipped"
	}

	a.db.Update(order)

	// 3. Retorna o pedido
	return order, nil
}