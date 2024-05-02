package events

import (
	"context"
	"github.com/IBM/sarama"
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"github.com/MuxiKeStack/be-feed/domain"
	"github.com/MuxiKeStack/be-feed/pkg/logger"
	"github.com/MuxiKeStack/be-feed/pkg/saramax"
	"github.com/MuxiKeStack/be-feed/service"
	"time"
)

const topicFeedEvent = "feed_event"

type FeedEvent struct {
	Type     feedv1.EventType
	Metadata map[string]string
}

type FeedEventConsumer struct {
	client sarama.Client
	l      logger.Logger
	svc    service.FeedService
}

func NewFeedEventConsumer(client sarama.Client, l logger.Logger, svc service.FeedService) *FeedEventConsumer {
	return &FeedEventConsumer{client: client, l: l, svc: svc}
}

func (f *FeedEventConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("feed-event-sync", f.client)
	if err != nil {
		return err
	}
	go func() {
		er := cg.Consume(context.Background(), []string{"feed_event"}, saramax.NewHandler(f.l, f.Consume))
		if er != nil {
			f.l.Error("退出了消费循环异常", logger.Error(er))
		}
	}()
	return nil
}

func (f *FeedEventConsumer) Consume(msg *sarama.ConsumerMessage, event FeedEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return f.svc.CreateFeedEvent(ctx, domain.FeedEvent{
		Type: event.Type,
		Ext:  event.Metadata,
	})
}
