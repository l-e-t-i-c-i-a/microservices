# Etapa de Compilação
FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

# Compila o Shipping usando o vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o shipping cmd/main.go

# Etapa Final
FROM scratch

WORKDIR /app
COPY --from=builder /app/shipping .

# Porta do Shipping
EXPOSE 3002

CMD ["./shipping"]