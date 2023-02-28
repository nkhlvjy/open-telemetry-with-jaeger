package main

import (
	pb "github.com/nkhlvjy/open-telemetry-with-jaeger/proto"
	"github.com/nkhlvjy/open-telemetry-with-jaeger/utils"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config := utils.ReadConfig()

	err = utils.SetGlobalTracer("checkout", config["JAEGER_ADDRESS"], config["JAEGER_PORT"])
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}

	channel, closeConn := utils.ConnectAmqp(config["RABBITMQ_USER"], config["RABBITMQ_PASS"], config["RABBITMQ_HOST"], config["RABBITMQ_PORT"])
	defer closeConn()

	lis, err := net.Listen("tcp", config["GRPC_ADDRESS"])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	pb.RegisterCheckoutServer(s, &server{channel: channel})

	log.Printf("GRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedCheckoutServer
	channel *amqp.Channel
}
