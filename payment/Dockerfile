# Atualizamos para 1.23 (ou 1.25) para ser compatível com seu projeto
FROM golang:1.24 AS builder

WORKDIR /usr/src/app

# Copia tudo (incluindo a pasta 'vendor' que criamos com o comando go mod vendor)
COPY . .

# IMPORTANTE: Adicionamos '-mod=vendor' para usar os arquivos locais do replace
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o payment ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/payment ./payment
EXPOSE 3001
CMD ["./payment"]