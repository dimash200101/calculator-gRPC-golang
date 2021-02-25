FROM golang

ADD proto/calc.pb.go /go/src/github.com/dimash200101/calculator-gRPC-golang/proto/
ADD server/main.go /go/src/github.com/dimash200101/calculator-gRPC-golang/server/

RUN go get -u golang.org/x/net/context
RUN go get -u google.golang.org/grpc
RUN go get -u google.golang.org/grpc/reflection

RUN go install github.com/dimash200101/calculator-gRPC-golang/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 4001
