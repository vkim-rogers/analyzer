FROM golang as builder

RUN go get github.com/gorilla/mux

WORKDIR /go/src/

COPY src .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o analyzer main/main.go

FROM ubuntu:18.04
COPY --from=builder /go/src/analyzer /bin/

EXPOSE 12345

CMD ["analyzer"]