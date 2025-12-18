# client.py

import grpc
import order_pb2
import order_pb2_grpc

def enviar_pedido(stub, nome_teste, customer_id, itens):
    """
    Função auxiliar para enviar o pedido e imprimir o resultado formatado.
    """
    print(f"🔵 EXECUTANDO: {nome_teste}")
    
    # Monta a requisição
    request = order_pb2.CreateOrderRequest(
        costumer_id=customer_id,
        order_items=itens
    )

    try:
        # Tenta enviar para o microsserviço Order
        response = stub.Create(request)
        print(f"✅ SUCESSO! Pedido criado com ID: {response.order_id}")
        print("   -> Verifique no banco: Status deve ser 'Paid'")
    
    except grpc.RpcError as e:
        # Captura o erro retornado pelo servidor
        print(f"❌ ERRO RECEBIDO (Status gRPC: {e.code()})")
        print(f"   Mensagem: {e.details()}")
        
        # Dicas do que verificar baseadas na mensagem
        if "exceed 50" in e.details():
             print("   -> Comportamento esperado para excesso de itens (Não salva no banco).")
        elif "Payment over 1000" in e.details():
             print("   -> Comportamento esperado para valor alto (Verifique no banco: Status deve ser 'Canceled').")
        else:
             print("   -> Erro não esperado.")
    
    print("-" * 40 + "\n")

def run():
    # Conectar ao servidor gRPC na porta 3000 (Order Service)
    print("🔌 Conectando ao servidor gRPC...")
    channel = grpc.insecure_channel('localhost:3000')
    stub = order_pb2_grpc.OrderStub(channel)
    print("-" * 40 + "\n")

    # --- CENÁRIO 1: Pedido Válido (Happy Path) ---
    # Qtd < 50 e Preço < 1000
    item_valido = order_pb2.OrderItem(
        product_code="CANETA",
        unit_price=10.0,
        quantity=5
    )
    enviar_pedido(stub, "Teste 1: Pedido Válido (5 itens, R$ 50)", 101, [item_valido])


    # --- CENÁRIO 2: Erro de Quantidade (Regra do Order) ---
    # Qtd > 50. Deve falhar ANTES de chamar o pagamento.
    item_muitos = order_pb2.OrderItem(
        product_code="CLIPES",
        unit_price=1.0,
        quantity=51 # Estourando o limite
    )
    enviar_pedido(stub, "Teste 2: Quantidade Exagerada (> 50 itens)", 102, [item_muitos])


    # --- CENÁRIO 3: Erro de Pagamento (Regra do Payment) ---
    # Preço > 1000. Deve salvar como Canceled.
    item_caro = order_pb2.OrderItem(
        product_code="NOTEBOOK",
        unit_price=1500.0, # Estourando o limite de preço
        quantity=1
    )
    enviar_pedido(stub, "Teste 3: Preço Alto (> R$ 1000)", 103, [item_caro])

if __name__ == '__main__':
    run()