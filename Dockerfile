FROM golang:1.17 as builder

# 使用不同的构建参数来选择不同微服务容器
ARG image=Api

# 设置工作目录
WORKDIR /go/src/go-study

# 复制go.mod
COPY go.mod go.mod
COPY go.sum go.sum

#使用cache下载
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct  && \
    go mod download

# 复制项目
COPY . .

# 构建应用
RUN go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go build -o /go/bin/${image} ./Container/${image}/main.go


FROM alpine:latest

WORKDIR /app
COPY --from=builder /go/bin/${image} .
COPY --from=builder /go/src/go-study/Config ./Config

#todo 提权运行构建文件
#ENV image=${image}
#RUN chmod +x ${image}
#RUN echo "./${image}"

# 端口
EXPOSE 8080