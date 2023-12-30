# ChatChit-backend

## Prerequisites
- Docker


## Build and run docker compose

```bash
docker compose -f build/docker-compose.yml up -d
```

## Or
## Setup Development Environment

```bash
docker run -d -p 6379:6379 --name my-redis redis
```

## Run

```bash
go run cmd/main.go
```