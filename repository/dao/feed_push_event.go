package dao

import (
	"context"
	"gorm.io/gorm"
)

type FeedPushEventDAO interface {
	CreateEvent(ctx context.Context, event FeedPushEvent) error
	FindFeedEvents(ctx context.Context, uid int64, lastTime int64, limit int64) ([]FeedPushEvent, error)
}

type GORMFeedPushEventDAO struct {
	db *gorm.DB
}

func NewGORMFeedPushEventDAO(db *gorm.DB) FeedPushEventDAO {
	return &GORMFeedPushEventDAO{db: db}
}

func (dao *GORMFeedPushEventDAO) FindFeedEvents(ctx context.Context, uid int64, lastTime int64, limit int64) ([]FeedPushEvent, error) {
	var events []FeedPushEvent
	err := dao.db.WithContext(ctx).
		Where("uid = ? and ctime > ?", uid, lastTime).
		Order("ctime asc").
		Limit(int(limit)).
		Find(&events).Error
	return events, err
}

func (dao *GORMFeedPushEventDAO) CreateEvent(ctx context.Context, event FeedPushEvent) error {
	return dao.db.WithContext(ctx).Create(&event).Error
}

type FeedPushEvent struct {
	Id      int64  `gorm:"primaryKey,autoIncrement"`
	UID     int64  `gorm:"column:uid;type:int(11);not null;"`
	Type    int32  `gorm:"column:type;type:varchar(255);comment:类型"`
	Content string `gorm:"column:content;type:text;"`
	// 发生时间
	Ctime int64 `gorm:"column:ctime;comment:发生时间"`
}
