FROM golang:alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add --no-cache git
ENV GOPROXY https://goproxy.cn

WORKDIR /app
COPY . .

RUN go build -v ./...
EXPOSE 60012

CMD ["/app/webim"]
