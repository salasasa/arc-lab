# --- 第一阶段：编译环境 (Builder) ---
# 使用与你本地一致的 Go 1.24 版本，alpine 版体积更小
FROM golang:1.24-alpine AS builder

# 1. 设置环境变量
# GO111MODULE=on: 开启模块支持
# GOPROXY: 使用七牛云代理加速国内依赖下载
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 2. 安装编译依赖
# tzdata 用于处理时区，ca-certificates 用于程序访问 HTTPS 外部接口
RUN apk add --no-cache tzdata ca-certificates

# 3. 设置工作目录
WORKDIR /app

# 4. 缓存依赖 (优化构建速度的关键)
# 只要 go.mod/sum 没变，后续构建会直接跳过 go mod download
COPY go.mod go.sum ./
RUN go mod download

# 5. 拷贝源码并编译
COPY . .
# CGO_ENABLED=0: 静态编译，保证在精简的 alpine 镜像中也能运行
# -ldflags="-s -w": 移除调试符号和符号表，极致压缩二进制体积
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gateway gateway.go


# --- 第二阶段：运行环境 (Final) ---
FROM alpine:latest

# 1. 准备运行时必要的环境
RUN apk --no-cache add ca-certificates tzdata
# 设置容器时区为北京时间
ENV TZ=Asia/Shanghai

WORKDIR /app

# 2. 从 builder 阶段拷贝编译好的程序和配置文件
# 这种方式保证了最终镜像不包含源代码，安全且轻量
COPY --from=builder /app/gateway /app/gateway
COPY --from=builder /app/etc/gateway-api.yaml /app/etc/gateway-api.yaml

# 3. 声明端口
EXPOSE 8888

# 4. 启动程序
# 必须使用 -f 参数指定容器内配置文件的路径
CMD ["./gateway", "-f", "etc/gateway-api.yaml"]