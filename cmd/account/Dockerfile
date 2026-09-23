FROM golang:1.26-alpine AS builder

WORKDIR /build

#hello devops or random developer. 
#i can write context: cmd/account/ in docker-compose
#but i can`t

COPY go.work ./
COPY db/ ./db/
COPY cmd/account/go.mod cmd/account/go.sum ./cmd/account/

RUN go mod download

COPY db/ ./db/
COPY cmd/account/ ./cmd/account/

WORKDIR /build/cmd/account

RUN CGO_ENABLED=0 GOWORK=/build/go.work go build -o account .

#==-==-==-==-==-==-==-==-==-==-==-==-==-==

FROM scratch

WORKDIR /app

COPY --from=builder /build/cmd/account/account .

CMD ["./account"]
