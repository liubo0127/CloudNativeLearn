# 一阶段编译
FROM golang:1.18.4 as build
WORKDIR /
COPY main.go .
ENV CGO_ENABLED=0 GO111MODULE=off GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct
RUN go get github.com/golang/glog
RUN go build -o httpserver .

# 二阶段构建镜像
FROM busybox
COPY --from=build /httpserver /
# 容器内使用的端口
EXPOSE 80
ENTRYPOINT ["/httpserver"]
CMD ["-log_dir", "/var/log/"]