package shipping_adapter

import (
	"context"
	"fmt"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	
	// Importa o proto do Shipping (que você gerou antes)
	"github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping"
	
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	// Mudou de PaymentClient para ShippingClient
	shipping shipping.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption

	// --- LÓGICA DE RETRY (MANTIDA IGUAL AO PAYMENT) ---
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1 * time.Second)),
		grpc_retry.WithMax(5),
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
	}

	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(shippingServiceUrl, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to shipping service: %v", err)
	}

	// Inicializa o cliente do Shipping
	client := shipping.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) ShipOrder(order domain.Order) error {
	// LÓGICA DE TIMEOUT (MANTIDA IGUAL)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// --- CONVERSÃO DE DADOS (DOMÍNIO -> PROTO) ---
	// O Shipping precisa da lista de itens para calcular o prazo.
	// Vamos converter os itens do domínio Order para o formato do Proto Shipping.
	var protoItems []*shipping.ShippingItem
	for _, item := range order.OrderItems {
		protoItems = append(protoItems, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	// CHAMADA gRPC AO SHIPPING
	_, err := a.shipping.Create(ctx, &shipping.CreateShippingRequest{
		OrderId: order.ID,
		Items:   protoItems, // Passamos a lista convertida
	})

	// TRATAMENTO DE ERRO
	if err != nil {
		if status.Code(err) == codes.DeadlineExceeded {
			log.Printf("TIMEOUT EXCEDIDO: O serviço de Shipping demorou mais de 2s para o pedido %d", order.ID)
		}
		return err
	}

	return nil
}