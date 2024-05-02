package repository

import (
	"context"
	"encoding/json"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/repository/dao"
	"github.com/ecodeclub/ekit/slice"
	"time"
)

type FeedEventRepository interface {
	CreatePushEvent(ctx context.Context, event domain.FeedEvent) error
	FindPushFeedEvents(ctx context.Context, uid int64, lastTime int64, limit int64) ([]domain.FeedEvent, error)
}

type feedEventRepository struct {
	pushDAO dao.FeedPushEventDAO
}

func NewFeedEventRepository(pushDAO dao.FeedPushEventDAO) FeedEventRepository {
	return &feedEventRepository{pushDAO: pushDAO}
}

func (repo *feedEventRepository) FindPushFeedEvents(ctx context.Context, uid int64, lastTime int64, limit int64) ([]domain.FeedEvent, error) {
	events, err := repo.pushDAO.FindFeedEvents(ctx, uid, lastTime, limit)
	return slice.Map(events, func(idx int, src dao.FeedPushEvent) domain.FeedEvent {
		return repo.convertToPushEventDomain(src)
	}), err
}

func (repo *feedEventRepository) CreatePushEvent(ctx context.Context, event domain.FeedEvent) error {
	return repo.pushDAO.CreateEvent(ctx, repo.convertToPushEventEntity(event))
}

func (repo *feedEventRepository) convertToPushEventEntity(event domain.FeedEvent) dao.FeedPushEvent {
	val, _ := json.Marshal(event.Ext)
	return dao.FeedPushEvent{
		Id:      event.ID,
		UID:     event.Uid,
		Type:    int32(event.Type),
		Content: string(val),
		Ctime:   event.Ctime.UnixMilli(),
	}
}

func (repo *feedEventRepository) convertToPushEventDomain(event dao.FeedPushEvent) domain.FeedEvent {
	var ext map[string]string
	_ = json.Unmarshal([]byte(event.Content), &ext)
	return domain.FeedEvent{
		ID:    event.Id,
		Uid:   event.UID,
		Type:  feedv1.EventType(event.Type),
		Ctime: time.UnixMilli(event.Ctime),
		Ext:   ext,
	}
}
