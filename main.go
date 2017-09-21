//go:generate protoc -I calendar --go_out=plugins=grpc:calendar calendar/calendar.proto

package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/integraal/chat-ops-calendar/calendar"
	"github.com/joho/godotenv"
	"github.com/caarlos0/env"
	"fmt"
	"strings"
	"golang.org/x/net/context"
)

type Server struct {
	BindAddress string `env:"GRPC_BIND" envDefault:":5001"`
	BaseUrl     string `env:"BASE_URL,required"`
}

func getServer() *Server {
	godotenv.Load()
	server := &Server{}
	err := env.Parse(server)
	if err != nil {
		panic(err)
	}
	return server
}

func (s *Server) getCalendarAddress(email string) string {
	return fmt.Sprintf(strings.TrimRight(s.BaseUrl, "/")+"/home/%s/Calendar.json", email)
}

type Parser struct {}

func (c *Parser) GetEvents() []*calendar.Event {
	// TODO: instantiate slice of calendar.Event
	return nil
}

func (c *Server) getParser(email string, start string, end string) (*Parser, error) {
	// TODO: check whether calendar is available, otherwise return error
	return &Parser{}, nil
}

func (s *Server) GetEvents(ctx context.Context, in *calendar.EventRequest) (*calendar.EventCollection, error) {
	parser, err := s.getParser(in.Email, in.Start, in.End)
	if err != nil {
		return nil, err
	}
	return &calendar.EventCollection{
		Items: parser.GetEvents(),
	}, nil
}

func main() {
	server := getServer()
	lis, err := net.Listen("tcp", server.BindAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Started listening " + server.BindAddress)
	rpc := grpc.NewServer()
	calendar.RegisterCalendarServer(rpc, server)
	// Register reflection service on gRPC server.
	reflection.Register(rpc)
	if err := rpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
