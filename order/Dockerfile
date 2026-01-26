# Etapa de Compilação
FROM golang:1.24 AS builder

WORKDIR /app

# Copia todo o código, INCLUINDO a pasta 'vendor' gerada no Passo 1
COPY . .

# Compila usando '-mod=vendor' (Usa os arquivos locais, não tenta baixar de novo)
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o order cmd/main.go

# Etapa Final (Imagem Leve)
FROM scratch

WORKDIR /app
COPY --from=builder /app/order .

# Porta do Order
EXPOSE 3000

CMD ["./order"]