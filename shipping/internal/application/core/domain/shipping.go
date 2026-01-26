package domain

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	OrderID      int64
	Items        []ShippingItem
	DeliveryDays int32
}

func NewShipping(orderId int64, items []ShippingItem) Shipping {
	// 1. Cria o objeto básico
	shipping := Shipping{
		OrderID: orderId,
		Items:   items,
	}
	
	// 2. Chama a lógica de cálculo imediatamente
	shipping.CalculateDelivery()

	return shipping
}

// Lógica de Negócio: Prazo mínimo 1 dia + 1 dia a cada 5 unidades
func (s *Shipping) CalculateDelivery() {
	totalQty := int32(0)
	for _, item := range s.Items {
		totalQty += item.Quantity
	}
	// Divisão de inteiros: 9/5 = 1. 
	s.DeliveryDays = 1 + (totalQty / 5)
}