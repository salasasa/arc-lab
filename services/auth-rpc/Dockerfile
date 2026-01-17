# 阶段一：编译
FROM golang:1.24.10-alpine AS builder
WORKDIR /app

# 安装必要的工具，alpine 镜像很小，有时需要补齐构建工具
RUN apk add --no-cache gcc musl-dev

# 配置代理（国内 N100 环境建议加上，构建更快）
ENV GOPROXY=https://goproxy.cn,direct

# 拷贝 go.mod
COPY go.mod go.sum ./
RUN go mod download

# 拷贝代码并编译
COPY . .
RUN go build -ldflags="-s -w" -o auth-rpc auth.go

# 阶段二：运行（使用最轻量的镜像）
FROM alpine:latest
WORKDIR /app

# 拷贝二进制文件和配置文件
COPY --from=builder /app/auth-rpc .
COPY --from=builder /app/etc/auth.yaml ./etc/auth.yaml

# 暴露端口
EXPOSE 8080

CMD ["./auth-rpc", "-f", "etc/auth.yaml"]