package payment_adapter

import (
	"context"
	"fmt"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/l-e-t-i-c-i-a/microservices-proto/golang/payment"
	"github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption

	// --- 1.2 CONFIGURAÇÃO DO RETRY (NOVAS TENTATIVAS) ---
	
	// Define as regras de retransmissão:
	// 1. BackoffLinear(1s): Espera 1s, 2s, 3s... a cada tentativa.
	// 2. WithMax(5): Tenta no máximo 5 vezes.
	// 3. WithCodes: Só tenta de novo se o erro for Unavailable (servidor caiu) ou ResourceExhausted.
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1 * time.Second)),
		grpc_retry.WithMax(5),
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
	}

	// Adiciona o "Interceptor" na lista de opções do gRPC.
	// O interceptor funciona como um "filtro" que captura a chamada e aplica a lógica de retry.
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)))

	// Configuração de segurança (insecure)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Cria a conexão (Dial) passando todas as opções (opts)
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to payment service: %v", err)
	}

	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	// context.WithTimeout cria um contexto que é cancelado automaticamente após 2 segundos.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// defer cancel() é crucial! Ele garante que os recursos do contexto sejam limpos
	// assim que a função Charge terminar, seja por sucesso ou erro.
	defer cancel()

	// CHAMADA COM O CONTEXTO
	// Passamos 'ctx' em vez de context.Background()
	_, err := a.payment.Create(ctx, &payment.CreatePaymentRequest{
		UserId: order.CustomerID,
		OrderId: order.ID,
		TotalPrice: order.TotalPrice(),
	})

	// TRATAMENTO DO ERRO DE TIMEOUT
	if err != nil {
		// status.Code(err) extrai o código gRPC do erro
		if status.Code(err) == codes.DeadlineExceeded {
			// Log solicitado na prática
			log.Printf("TIMEOUT EXCEDIDO: A chamada ao serviço de Payment demorou mais de 2 segundos para o pedido %d", order.ID)
		}
		return err
	}


	return nil
}