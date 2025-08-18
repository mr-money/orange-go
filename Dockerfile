# 多段构建 builder构建二进制文件
FROM golang:1.23 as builder

# 使用不同的构建参数来选择不同微服务容器
ARG image=Api

# 设置工作目录
WORKDIR /go/src/orange-go

# 复制依赖描述文件
COPY go.mod go.mod
COPY go.sum go.sum

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download

# 复制源码
COPY . .

# 构建应用
RUN go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go build -o /go/bin/app ./Container/${image}/main.go

# 服务容器运行
FROM alpine:latest

# 时区
ENV TZ=Asia/Shanghai
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /go/bin/app .
COPY --from=builder /go/src/orange-go/Config ./Config

ENTRYPOINT ["./app"]

# 端口
EXPOSE 8080