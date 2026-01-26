package main

import (
	"log"

	"github.com/l-e-t-i-c-i-a/microservices/order/config"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/adapters/db"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/adapters/payment"

	shipping_adapter "github.com/l-e-t-i-c-i-a/microservices/order/internal/adapters/shipping"

	//"github.com/l-e-t-i-c-i-a/microservices/order/internal/adapters/rest"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/adapters/grpc"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/api"
)

func main() {
	// Banco de Dados
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	// Payment
	paymentAdapter, err := payment_adapter.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	// Inicializa o adaptador do Shipping (NOVO)
	shippingAdapter, err := shipping_adapter.NewAdapter(config.GetShippingServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize shipping stub. Error: %v", err)
	}

	// Application
	application := api.NewApplication(dbAdapter, paymentAdapter, shippingAdapter)

	// Servidor gRPC
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	log.Println("Order Service is running...")
	grpcAdapter.Run()
}