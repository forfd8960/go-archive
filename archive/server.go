package archive

import (
	"context"

	pb "github.com/forfd8960/go-archive/pb"
)

type Server struct {}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetArchiveList(ctx context.Context, req *pb.GetArchiveListReq) (*pb.GetArchiveListRes, error) {
	return new(pb.GetArchiveListRes), nil
}
