package ports

import "github.com/l-e-t-i-c-i-a/microservices/order/internal/application/core/domain"

type ShippingPort interface {
    // ShipOrder recebe o pedido e chama o microsserviço de entrega
    ShipOrder(order domain.Order) error
}