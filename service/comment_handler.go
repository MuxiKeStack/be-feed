package service

import (
	"context"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/repository"
	"time"
)

type CommentFeedEventHandler struct {
	repo repository.FeedEventRepository
}

// Metadata: map[string]string{
// // 评论者
// "commentator": strconv.FormatInt(comment.Commentator.ID, 10),
// // 被评论者
// "recipient": strconv.FormatInt(comment.ReplyToUid, 10),
// // 资源发布者，可能与被评论者相同
// "bizPublisher": strconv.FormatInt(publisherId, 10),
// "biz":          comment.Biz.String(),
// "bizId":        strconv.FormatInt(comment.BizId, 10),
// "commentId":    strconv.FormatInt(comment.Id, 10),
// },
func (c *CommentFeedEventHandler) CreateFeedEvent(ctx context.Context, ext domain.ExtendFields) error {
	bizPublisher, err := ext.Get("bizPublisher").AsInt64()
	if err != nil {
		return err
	}
	recipient, err := ext.Get("recipient").AsInt64()
	if err != nil {
		return err
	}
	err = c.repo.CreatePushEvent(ctx, domain.FeedEvent{
		Uid:   bizPublisher,
		Type:  feedv1.EventType_Comment,
		Ctime: time.Now(),
		Ext:   ext,
	})
	if err != nil {
		return err
	}
	if recipient != bizPublisher {
		return c.repo.CreatePushEvent(ctx, domain.FeedEvent{
			Uid:   recipient,
			Type:  feedv1.EventType_Comment,
			Ctime: time.Now(),
			Ext:   ext,
		})
	}
	return nil
}
