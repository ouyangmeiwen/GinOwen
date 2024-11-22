# 使用 Go 官方镜像作为构建基础
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /app

# 仅复制依赖文件，优化缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制项目代码
COPY . .

# 构建可执行文件
RUN go build -o ginowen

# 使用轻量级镜像运行程序
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/ginowen .

CMD ["./ginowen"]
