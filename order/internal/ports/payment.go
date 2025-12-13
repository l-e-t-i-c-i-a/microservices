package ports

import "github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}