FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN go test -v .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fidibo main.go
RUN chmod +x fidibo

CMD ["/app/fidibo"]