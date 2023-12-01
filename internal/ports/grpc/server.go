package grpc_

import (
	"fmt"
	protos "github.com/nesrayr/protos/gen"
	"google.golang.org/grpc"
	"net"
	log "url_shortener/pkg/logging"
)

type Server struct {
	server *grpc.Server
	log    *log.Logger
}

func NewServer(srv protos.UrlShortenerServer, l log.Logger) *Server {
	server := grpc.NewServer()
	protos.RegisterUrlShortenerServer(server, srv)

	return &Server{server: server, log: &l}
}

func (s *Server) Run(port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		s.log.Error(err)
		return err
	}

	s.log.Info("grpc server successfully started")

	err = s.server.Serve(listener)
	if err != nil {
		s.log.Error(err)
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.log.Info("stopping grpc server")

	s.server.GracefulStop()
}
