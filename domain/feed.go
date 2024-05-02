package domain

import (
	feedv1 "github.com/MuxiKeStack/be-api/gen/proto/feed/v1"
	"time"
)

type FeedEvent struct {
	ID int64
	// 以 A 发表了一篇文章为例
	// 如果是 Pull Event，也就是拉模型，那么 Uid 是 A 的id
	// 如果是 Push Event，也就是推模型，那么 Uid 是 A 的某个粉丝的 id
	Uid   int64
	Type  feedv1.EventType
	Ctime time.Time
	Ext   ExtendFields
}
