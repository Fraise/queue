package server

import (
	"errors"
	"google.golang.org/grpc"
	server "server/server/grpc"
)

type Server struct {
	server.UnimplementedQueueServer
}

func (s *Server) Publish(stream grpc.BidiStreamingServer[server.PublishRequest, server.PublishResponse]) error {
	return errors.New("unimplemented")
}
