# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.21.0-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ------------------------------------------------------
# デプロイ用コンテナ
FROM debian:bullseye-slim as deploy

RUN opt-get update

COPY --from=deploy-builder /app/app .
CMD ["./app"]

# ------------------------------------------------------
# ローカル開発用コンテナ（ホットリロード）
FROM golang:1.21.0 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
