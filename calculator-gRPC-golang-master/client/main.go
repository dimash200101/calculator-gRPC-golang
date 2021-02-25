package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/dimash200101/calculator-gRPC-golang/proto"
)

const (
	address = "localhost:4001"
)

func argParser(n1 string, n2 string) (int32, int32) {
	N1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot parse arge[1]: %s", err)
	}
	N2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Cannot parse arge[2]: %s", err)
	}
	return int32(N1), int32(N2)
}

func main() {
	if len(os.Args) != 3 {
            log.Fatalf("2 numbers expected: n1 n2")
	}

	n1, n2 := argParser(os.Args[1], os.Args[2])

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addResult, err := client.Add(ctx, &pb.AddRequest{int32(n1), int32(n2)})
	if err != nil {
		log.Fatalf("Adding error: %s", err)
	}
	log.Printf("(%d + %d + %d + %d = %d)  / 4", n1, n2, n3, n4 addResult.N1)
