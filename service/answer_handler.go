package service

import (
	"context"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/repository"
	"time"
)

type AnswerFeedEventHandler struct {
	repo repository.FeedEventRepository
}

// "answerer":   strconv.FormatInt(answer.PublisherId, 10),
// "questioner": strconv.FormatInt(res.GetQuestion().GetQuestionerId(), 10),
// "questionId": strconv.FormatInt(answer.QuestionId, 10),
// "answerId":   strconv.FormatInt(answer.Id, 10),
func (a *AnswerFeedEventHandler) CreateFeedEvent(ctx context.Context, ext domain.ExtendFields) error {
	questioner, err := ext.Get("questioner").AsInt64()
	if err != nil {
		return err
	}
	return a.repo.CreatePushEvent(ctx, domain.FeedEvent{
		Uid:   questioner,
		Type:  feedv1.EventType_Answer,
		Ctime: time.Now(),
		Ext:   ext,
	})
}
