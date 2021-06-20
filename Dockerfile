FROM golang:1.16 as build

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE=on
WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/app /
COPY --from=build /go/release/vendor /vendor

CMD ["/app"]
