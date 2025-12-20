## Passos para testar
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

#### 4. Dentro de microservices/order:
- Abrir um terminal
- Rodar:
```powershell

# Configura variáveis (incluindo a nova URL do Payment)
$env:DB_DRIVER="mysql"
$env:DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:3306)/order?parseTime=true"
$env:APPLICATION_PORT="3000"
$env:ENV="development"
$env:PAYMENT_SERVICE_URL="localhost:3001"

# Roda o serviço
go run cmd/main.go

```

#### 5. Dentro de microservices/client:
- Abrir um terminal
- Rodar:
```powershell

py client.py

```

##### Se aparecer "Pedido criado com ID:", funcionou!

#### 6. Para ver os status dos pedidos
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
```
