FROM golang:1.17.5 AS build-env
ENV GOPROXY https://goproxy.cn
ADD . /go/src/app
WORKDIR /go/src/app
RUN go mod tidy
RUN cd cmd && GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-extldflags "-static"' -o /go/src/app/app-server /go/src/app/cmd/main.go

FROM registry.cn-hangzhou.aliyuncs.com/coolops/ubuntu:22.04
ENV TZ=Asia/Shanghai
COPY --from=build-env /go/src/app/app-server /opt/app-server
WORKDIR /opt
EXPOSE 80
CMD [ "./app-server" ]
