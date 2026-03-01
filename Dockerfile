FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copiar solo go.mod y go.sum primero para aprovechar caché
COPY go.mod go.sum ./

# Descargar módulos en una capa separada (caché)
RUN go mod download && go mod verify

# Copiar el resto del código
COPY . .

# Compilar con optimizaciones
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -a -installsuffix cgo \
    -o app ./cmd/server/main.go

# Stage final: alpine optimizado
FROM alpine:3.20

WORKDIR /app

# Instalar solo lo necesario para health checks
RUN apk add --no-cache ca-certificates

# Copiar el binario compilado
COPY --from=builder /app/app .

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/send || exit 1

CMD ["./app"]
