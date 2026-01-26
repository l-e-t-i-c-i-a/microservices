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
        # ATUALIZAÇÃO: Agora o fluxo completo vai até o Shipping
        print("   -> EXPECTATIVA: Status no banco deve ser 'Shipped' (Enviado)")
    
    except grpc.RpcError as e:
        # Captura o erro retornado pelo servidor
        print(f"❌ ERRO RECEBIDO (Status gRPC: {e.code()})")
        print(f"   Mensagem: {e.details()}")
        
        # Dicas do que verificar baseadas na mensagem
        error_msg = e.details() if e.details() else ""
        
        if "exceed 50" in error_msg:
             print("   -> ✅ OK! Bloqueio de quantidade funcionou.")
        
        elif "Payment over 1000" in error_msg:
             print("   -> ✅ OK! Bloqueio de pagamento funcionou.")
             print("   -> EXPECTATIVA: Status no banco deve ser 'Canceled'")
        
        elif "product not found" in error_msg:
             print("   -> ✅ OK! Validação de Estoque funcionou (Requisito 1.2).")
        
        elif "connectex" in error_msg or "unavailable" in error_msg.lower():
             print("   -> ⚠️  ALERTA: Parece que um dos microsserviços está desligado.")
        
        else:
             print("   -> ⚠️  Erro não esperado ou Timeout.")
    
    print("-" * 50 + "\n")

def run():
    # Conectar ao servidor gRPC na porta 3000 (Order Service)
    print("🔌 Conectando ao servidor gRPC (localhost:3000)...")
    channel = grpc.insecure_channel('localhost:3000')
    stub = order_pb2_grpc.OrderStub(channel)
    print("-" * 50 + "\n")

    # --- CENÁRIO 1: Pedido Válido (Happy Path Completo) ---
    # Produto existe, Qtd OK, Preço OK.
    # Deve passar pelo Order -> DB -> Payment -> Shipping
    item_valido = order_pb2.OrderItem(
        product_code="CANETA", # Este item foi inserido no seed do DB
        unit_price=10.0,
        quantity=5
    )
    enviar_pedido(stub, "Teste 1: Happy Path (Tudo Certo)", 101, [item_valido])


    # --- CENÁRIO 2: Produto Inexistente (Teste de Estoque) ---
    # Requisito 1.2: Deve falhar antes de salvar
    item_fantasma = order_pb2.OrderItem(
        product_code="BOLA_QUADRADA", # Item que NÃO existe no banco
        unit_price=50.0,
        quantity=1
    )
    enviar_pedido(stub, "Teste 2: Produto Inexistente (Validação de Estoque)", 102, [item_fantasma])


    # --- CENÁRIO 3: Erro de Quantidade ---
    # Regra de Negócio: Qtd > 50
    item_muitos = order_pb2.OrderItem(
        product_code="CLIPES", # Código irrelevante aqui, falha antes
        unit_price=1.0,
        quantity=51 
    )
    enviar_pedido(stub, "Teste 3: Quantidade Exagerada (> 50)", 103, [item_muitos])


    # --- CENÁRIO 4: Erro de Pagamento ---
    # Regra: Total > 1000. Salva como Canceled.
    item_caro = order_pb2.OrderItem(
        product_code="NOTEBOOK", # Este item existe no seed
        unit_price=1500.0, 
        quantity=1
    )
    enviar_pedido(stub, "Teste 4: Preço Alto (> R$ 1000)", 104, [item_caro])

if __name__ == '__main__':
    run()