FROM golang:1.25 AS builder
WORKDIR /app

# Cette fois, la copie inclut le go.mod car on est à la racine !
COPY . .

# IMPORTANT : Remplacez ici le chemin vers votre dossier barrgraph
# Par exemple : RUN CGO_ENABLED=0 GOOS=linux go build -o barrgraph ./dsme/barrgraph/go/
RUN CGO_ENABLED=0 GOOS=linux go build -o barrgraph ./dsme/barrgraph/go/cmd/barrgraph

FROM scratch
COPY --from=builder /app/barrgraph /barrgraph
EXPOSE 8080
ENTRYPOINT ["/barrgraph"]