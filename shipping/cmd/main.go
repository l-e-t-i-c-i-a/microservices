package main

import (
	"log"

	"github.com/l-e-t-i-c-i-a/microservices/shipping/config"
	"github.com/l-e-t-i-c-i-a/microservices/shipping/internal/adapters/grpc"
	"github.com/l-e-t-i-c-i-a/microservices/shipping/internal/application/core/api"
)

func main() {
	// 1. Inicializa a Aplicação (Core)
	// Diferente do Payment, aqui não precisamos passar adaptador de banco de dados
	application := api.NewApplication()

	// 2. Inicializa o Adaptador gRPC (Servidor)
	// Passamos a aplicação e a porta configurada
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	// 3. Roda o servidor
	log.Printf("Starting Shipping Service on port %d...", config.GetApplicationPort())
	grpcAdapter.Run()
}