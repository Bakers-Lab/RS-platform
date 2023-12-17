# 使用一个Go构建镜像
FROM golang:1.20 as builder

# 设置工作目录
WORKDIR /app

# 复制Go模块文件
COPY go.mod go.sum ./

# 下载依赖
RUN go env -w GO111MODULE=on
RUN go mod tidy
# RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN go clean -modcache && CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用一个轻量级的基础镜像
FROM alpine:latest  

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制构建的二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]
