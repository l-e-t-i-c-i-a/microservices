# 📦 Projeto Microservices gRPC (Order, Payment, Shipping)

Este projeto implementa uma arquitetura de microsserviços para um E-commerce utilizando **Go**, **gRPC**, **MySQL** e **Docker**.

O sistema é composto por:
- **Order Service:** Gerencia pedidos.
- **Payment Service:** Processa pagamentos.
- **Shipping Service:** Calcula prazos de entrega.

---

## 🛠️ Pré-requisitos
* **Go** 1.23+
* **Docker** e **Docker Compose**
* **Python** 3.x (para rodar o script de teste do cliente)
* Bibliotecas Python: `pip install grpcio grpcio-tools`

---
## Passos para testar

## 🚀 Opção 1: Rodar com Docker Compose (Recomendado)
Esta é a maneira mais simples de executar, pois sobe o banco de dados e os 3 serviços automaticamente com todas as dependências já configuradas.

### Subir a Aplicação
Na pasta raiz microservices, execute: docker-compose up --build
Aguarde até ver logs indicando que os serviços "Order", "Payment" e "Shipping" iniciaram nas portas 3000, 3001 e 3002.

### Rodar o cliente
- Abrir um terminal  
- Rodar:
```powershell

py client.py

```


## 🚀 Opção 2: Rodar manualmente
#### 1. Baixar os repositórios microservices e microservices-proto e deixá-los numa mesma pasta

#### 2. Dentro de microservices:
- Abrir um terminal  
- Rodar:
```powershell

docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=minhasenha -v "${PWD}/init.sql:/docker-entrypoint-initdb.d/init.sql" mysql

```
#### 3. Dentro de microservices/payment:
- Abrir um terminal
- Rodar:
```powershell
 
 # Configura variáveis
 $env:DB_DRIVER="mysql"
 $env:DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:3306)/payment"
 $env:APPLICATION_PORT="3001"
 $env:ENV="development"
 
 # Roda o serviço
 go run cmd/main.go
 
```

#### 4. Dentro de microservices/shipping:
- Abrir um terminal
- Rodar:
```powershell
 
 # Configura variáveis
 $env:APPLICATION_PORT="3002"
 $env:ENV="development"
 
 # Roda o serviço
 go run cmd/main.go
 
```

#### 5. Dentro de microservices/order:
- Abrir um terminal
- Rodar:
```powershell

# Configura variáveis (incluindo a nova URL do Payment)
$env:DB_DRIVER="mysql"
$env:DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:3306)/order?parseTime=true"
$env:APPLICATION_PORT="3000"
$env:ENV="development"
$env:PAYMENT_SERVICE_URL="localhost:3001"
$env:SHIPPING_SERVICE_URL="localhost:3002"

# Roda o serviço
go run cmd/main.go

```

#### 6. Dentro de microservices/client:
- Abrir um terminal
- Rodar:
```powershell

py client.py

```

#### 7. Para ver os status dos pedidos
- No terminal 1 onde rodou o Docker digite: docker ps (para ver o ID do container mysql)
- Rodar:
```powershell

# Troque 'ID_AQUI' pelo ID do seu container
docker exec -it ID_AQUI mysql -u root -pminhasenha

```

- Quando o terminal mudar para mysql>, rode o SQL:
```SQL
USE `order`;
SELECT * FROM orders;
SELECT * FROM products;
```