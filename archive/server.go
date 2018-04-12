package archive

import (
	"context"

	pb "github.com/forfd8960/go-archive/pb"
)

type Server struct {
	item Itemer
}

func NewServer(item Itemer) *Server {
	return &Server{item: item}
}

func (s *Server) GetArchiveList(ctx context.Context, req *pb.GetArchiveListReq) (*pb.GetArchiveListRes, error) {
	items, err := s.item.GetItemList(req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	count, err := s.item.GetItemCount()
	if err != nil {
		return nil, err
	}

	var pbItems []*pb.ArchiveItem
	for _, t := range items {
		pbItems = append(pbItems, &pb.ArchiveItem{
			FromHost:  t.FromHost,
			ItemName:  t.Name,
			ItemUrl:   t.URL,
			CollectAt: t.CollectAt.String(),
		})
	}

	return &pb.GetArchiveListRes{ArchiveItems: pbItems, TotalCount: count}, nil
}
