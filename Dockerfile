FROM golang:alpine3.20 AS builder

# Linux
LABEL maintainer="Abhinawa vps"
RUN apk update && apk add --no-cache git

# app
WORKDIR /app

COPY api-traceroot/go.mod api-traceroot/go.sum ./
RUN go get -u ./...
COPY api-traceroot .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM ubuntu:24.10

WORKDIR /root/
RUN apt-get update && apt-get install -y ca-certificates \
    && apt-get install -y software-properties-common \
    && add-apt-repository --yes --update ppa:ansible/ansible \
    && apt-get install -y ansible 
    
RUN apt-get install -y python3.12-venv \
    && python3 -m venv myenv \
    && myenv/bin/pip install pyvmomi \
    && myenv/bin/pip install requests

COPY --from=builder /app/main .
COPY /api-traceroot/.env .
COPY ansible /root

CMD ["./main"]