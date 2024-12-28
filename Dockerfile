FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o websocket cmd/main.go

# Crear una imagen mínima para ejecutar la aplicación
FROM debian:bookworm-slim

WORKDIR /app

# Copiar el binario compilado desde la etapa anterior
COPY --from=builder /app/websocket /app/websocket

# Exponer el puerto donde corre la aplicación
EXPOSE 8080

# Comando para iniciar la aplicación
CMD ["/app/websocket"]
