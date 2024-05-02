package service

import (
	"context"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/repository"
	"time"
)

type InviteFeedEventHandler struct {
	repo repository.FeedEventRepository
}

// Metadata: map[string]string{
// "inviter":    strconv.FormatInt(inviter, 10),
// "invitee":    strconv.FormatInt(invitee, 10),
// "biz":        question.Biz.String(), // 传出当前服务，则该枚举值变为易于理解的string
// "bizId":      strconv.FormatInt(question.BizId, 10),
// "questionId": strconv.FormatInt(question.Id, 10),
// },
func (i *InviteFeedEventHandler) CreateFeedEvent(ctx context.Context, ext domain.ExtendFields) error {
	// 不会邀请很多，推模型
	// invitee 受邀者是收件人
	// TODO inviter != invitee
	invitee, err := ext.Get("invitee").AsInt64()
	if err != nil {
		return err
	}
	return i.repo.CreatePushEvent(ctx, domain.FeedEvent{
		Uid:   invitee,
		Type:  feedv1.EventType_InviteToAnswer,
		Ctime: time.Now(),
		Ext:   ext,
	})
}
