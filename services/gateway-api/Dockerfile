FROM golang:1.24.10-alpine AS builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
ENV GOPROXY=https://goproxy.cn,direct

# 重点：从根目录上下文拷贝这两个模块，保持相对路径结构
COPY services/auth-rpc /app/auth-rpc
COPY services/gateway-api /app/gateway-api

# 进入网关目录进行构建
WORKDIR /app/gateway-api
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/gateway-api-bin gateway.go

FROM alpine:latest
WORKDIR /app
# 拷贝编译后的二进制文件
COPY --from=builder /app/gateway-api-bin .
# 拷贝配置文件
COPY --from=builder /app/gateway-api/etc/gateway-api.yaml ./etc/gateway-api.yaml

EXPOSE 8888
CMD ["./gateway-api-bin", "-f", "etc/gateway-api.yaml"]