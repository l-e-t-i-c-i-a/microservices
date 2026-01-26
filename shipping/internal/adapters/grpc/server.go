package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/l-e-t-i-c-i-a/microservices-proto/golang/shipping"
	"github.com/l-e-t-i-c-i-a/microservices/shipping/config"
	"github.com/l-e-t-i-c-i-a/microservices/shipping/internal/application/core/domain"
	"github.com/l-e-t-i-c-i-a/microservices/shipping/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	shipping.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
    
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}

func (a Adapter) Create(ctx context.Context, request *shipping.CreateShippingRequest) (*shipping.CreateShippingResponse, error) {
	log.Printf("🚢 SHIPPING: Recebido pedido de envio para Order ID: %d", request.OrderId)

	var items []domain.ShippingItem
	for _, i := range request.Items {
		items = append(items, domain.ShippingItem{ProductCode: i.ProductCode, Quantity: i.Quantity})
	}

	newShipping := domain.NewShipping(request.OrderId, items)
	result, err := a.api.CreateShipping(newShipping)
	if err != nil {
		log.Printf("❌ SHIPPING: Erro ao criar envio: %v", err)
		return nil, err
	}

	log.Printf("✅ SHIPPING: Envio processado. Prazo calculado: %d dias", result.DeliveryDays)

	return &shipping.CreateShippingResponse{
        ShippingId: 123, // Fake ID
        DeliveryDays: result.DeliveryDays,
    }, nil
}