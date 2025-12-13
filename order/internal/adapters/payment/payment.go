package payment_adapter

import (
	"context"
	//"log"

	"github.com/l-e-t-i-c-i-a/microservices-proto/golang/payment"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := payment.NewPaymentClient(conn) //inicializa o stub
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId: order.CustomerID,
		OrderId: order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}