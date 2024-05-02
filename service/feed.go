package service

import (
	"context"
	"fmt"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/repository"
)

type feedService struct {
	repo                repository.FeedEventRepository
	feedEventHandlerMap map[feedv1.EventType]FeedEventHandler
}

func NewFeedService(repo repository.FeedEventRepository) FeedService {
	return &feedService{
		repo: repo,
		feedEventHandlerMap: map[feedv1.EventType]FeedEventHandler{
			feedv1.EventType_InviteToAnswer: &InviteFeedEventHandler{repo: repo},
			feedv1.EventType_Support:        &SupportFeedEventHandler{repo: repo},
			feedv1.EventType_Comment:        &CommentFeedEventHandler{repo: repo},
			feedv1.EventType_Answer:         &AnswerFeedEventHandler{repo: repo},
		},
	}
}

func (f *feedService) FindFeedEvents(ctx context.Context, uid int64, lastTime int64, limit int64) ([]domain.FeedEvent, error) {
	return f.repo.FindPushFeedEvents(ctx, uid, lastTime, limit)
}

func (f *feedService) CreateFeedEvent(ctx context.Context, ft domain.FeedEvent) error {
	handler, exists := f.feedEventHandlerMap[ft.Type]
	if !exists {
		return fmt.Errorf("非法的feed事件类型：%s", ft.Type.String())
	}
	return handler.CreateFeedEvent(ctx, ft.Ext)
}
