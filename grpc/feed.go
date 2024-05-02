package grpc

import (
	"context"
	"encoding/json"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/service"
	"github.com/ecodeclub/ekit/slice"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type FeedServiceServer struct {
	feedv1.UnimplementedFeedServiceServer
	svc service.FeedService
}

func NewFeedServiceServer(svc service.FeedService) *FeedServiceServer {
	return &FeedServiceServer{svc: svc}
}

func (f *FeedServiceServer) FindFeedEvents(ctx context.Context, request *feedv1.FindFeedEventsRequest) (*feedv1.FindFeedEventsResponse, error) {
	events, err := f.svc.FindFeedEvents(ctx, request.GetUid(), request.GetLastTime(), request.GetLimit())
	return &feedv1.FindFeedEventsResponse{
		FeedEvents: slice.Map(events, func(idx int, src domain.FeedEvent) *feedv1.FeedEvent {
			return convertToV(src)
		}),
	}, err
}

func (f *FeedServiceServer) Register(server *grpc.Server) {
	feedv1.RegisterFeedServiceServer(server, f)
}

func convertToV(event domain.FeedEvent) *feedv1.FeedEvent {
	val, _ := json.Marshal(event)
	return &feedv1.FeedEvent{
		Id:      event.ID,
		Uid:     event.Uid,
		Type:    event.Type.String(),
		Content: string(val),
		Ctime:   event.Ctime.UnixMilli(),
	}
}
