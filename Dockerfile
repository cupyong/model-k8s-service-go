FROM docker.dm-ai.cn/smart-city/golang:1.19.2-alpine3.16

ENV TZ=Asia/Shanghai
WORKDIR /model-k8s-service-go
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go build -o main  main.go
EXPOSE 80
CMD ["./main"]
