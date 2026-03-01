# Étape 1 : On utilise une image avec Go pour compiler
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
# On compile explicitement pour Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o barrgraph .

# Étape 2 : On crée l'image finale ultra légère
FROM scratch
COPY --from=builder /app/barrgraph /barrgraph
EXPOSE 8080
ENTRYPOINT ["/barrgraph"]