package ports

import "github.com/l-e-t-i-c-i-a/microservices/shipping/internal/application/core/domain"

type APIPort interface {
	CreateShipping(shipping domain.Shipping) (domain.Shipping, error)
}