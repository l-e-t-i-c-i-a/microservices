# client.py

import grpc
import order_pb2
import order_pb2_grpc

def run():
    # Conectar ao servidor gRPC (ajuste o host/porta conforme necessário)
    channel = grpc.insecure_channel('localhost:3000')
    stub = order_pb2_grpc.OrderStub(channel)

    # Criar itens do pedido
    item1 = order_pb2.OrderItem(
        product_code="ABC123",
        unit_price=10.5,
        quantity=2
    )
    item2 = order_pb2.OrderItem(
        product_code="XYZ789",
        unit_price=20.0,
        quantity=1
    )

    # Criar a requisição
    request = order_pb2.CreateOrderRequest(
        costumer_id=42,
        order_items=[item1, item2],
        total_price=41.0
    )

    # Enviar a requisição
    response = stub.Create(request)

    print("Order created with ID:", response.order_id)

if __name__ == '__main__':
    run()