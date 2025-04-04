# 🔹 Etapa 1: Construcción del binario
FROM golang:1.22 AS builder

WORKDIR /app

# Copiar archivos y módulos de Go
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar el binario con optimización
RUN CGO_ENABLED=0 GOOS=linux go build -o webhook .

# 🔹 Etapa 2: Imagen final mínima
FROM alpine:latest  

WORKDIR /root/

# Copiar el binario desde la etapa anterior
COPY --from=builder /app/webhook .

# Exponer el puerto 5001 para el webhook
EXPOSE 8000

# Establecer la variable de entorno PORT
ENV PORT=8000

# Ejecutar el servicio
CMD ["./webhook"]


