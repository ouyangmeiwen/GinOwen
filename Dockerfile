# 使用 Go 官方镜像作为构建基础
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/golang:1.23.1 as builder

# 设置工作目录
WORKDIR /app

# 设置 Go proxy 为国内镜像，加速下载
ENV GOPROXY=https://goproxy.cn,direct

# 仅复制依赖文件，优化缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制项目代码
COPY . .

# 静态编译 Go 可执行文件，避免依赖 C 库
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -buildvcs=false -o ginowen

# 使用轻量级 Alpine 镜像运行程序
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/alpine:latest

# 安装依赖库（如果有必要）
# 如果你的应用需要使用特定的 C 库，可以安装相应的包
# RUN apk add --no-cache libc6-compat

# 设置工作目录
WORKDIR /root/

# 复制构建好的可执行文件
COPY --from=builder /app/ginowen .


# 复制配置文件 config.yaml
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/docs/swagger.json ./docs/swagger.json
COPY --from=builder /app/docs/swagger.yaml ./docs/swagger.yaml




# 确保文件具有执行权限
RUN chmod +x /root/ginowen

EXPOSE 9003

# 设置容器启动命令
CMD ["/root/ginowen"]

