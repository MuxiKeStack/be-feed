package service

import (
	"context"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
)

type FeedService interface {
	CreateFeedEvent(ctx context.Context, ft domain.FeedEvent) error
	FindFeedEvents(ctx context.Context, uid int64, lastTime int64, direction feedv1.Direction, limit int64) ([]domain.FeedEvent, error)
}

type FeedEventHandler interface {
	CreateFeedEvent(ctx context.Context, ext domain.ExtendFields) error
}
