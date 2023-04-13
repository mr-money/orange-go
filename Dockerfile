# https://zhuanlan.zhihu.com/p/269115851
FROM golang:1.17 as builder

# 使用不同的构建参数来选择不同的基础镜像
ARG image=Api

# 设置工作目录
WORKDIR /go/src/go-study

# 复制项目
COPY . .

# 构建应用
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod tidy && \
    go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go build -o /go/bin/${image} ./Container/${image}/main.go


FROM alpine:latest

WORKDIR /app
COPY --from=builder /go/bin/${image} .
COPY --from=builder /go/src/go-study/Config ./Config

# 端口
#EXPOSE 8080

#RUN go build -o go_study_api .
#WORKDIR /bin/go_study_api

#FROM golang:1.17-alpine
#LABEL image.authors="Mr_Money"

#COPY E:/Gopath/src/go-study/ ./src/go_study/
#RUN go build -o go_study_api ./src/go_study/Container/Api/main.go
#WORKDIR /bin/go_study_api

#RUN go env -w GO111MODULE=on && \
#    go env -w GOPROXY=https://goproxy.cn,direct && \
#    go mod tidy && \
#    go export GOARCH=amd64 && \
#    go export GOOS=linux && \
#    go build -o ../bin/go_study_api ./main.go && \
#    chmod +x go_study_api

CMD ["./${image}"]