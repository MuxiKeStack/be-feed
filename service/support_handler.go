package service

import (
	"context"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/repository"
	"time"
)

type SupportFeedEventHandler struct {
	repo repository.FeedEventRepository
}

// Metadata: map[string]string{
// "supporter": strconv.FormatInt(ubs.Uid, 10),
// "supported": strconv.FormatInt(supported, 10),
// "biz":       ubs.Biz.String(),
// "bizId":     strconv.FormatInt(ubs.BizId, 10),
// },
func (s *SupportFeedEventHandler) CreateFeedEvent(ctx context.Context, ext domain.ExtendFields) error {
	supported, err := ext.Get("supported").AsInt64()
	if err != nil {
		return err
	}
	return s.repo.CreatePushEvent(ctx, domain.FeedEvent{
		Uid:   supported,
		Type:  feedv1.EventType_Support,
		Ctime: time.Now(),
		Ext:   ext,
	})
}
