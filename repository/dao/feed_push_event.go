package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type FeedPushEventDAO interface {
	CreateEvent(ctx context.Context, event FeedPushEvent) error
	FindFeedEvents(ctx context.Context, uid int64, lastTime int64, direction int32, limit int64) ([]FeedPushEvent, error)
}

type GORMFeedPushEventDAO struct {
	db *gorm.DB
}

func NewGORMFeedPushEventDAO(db *gorm.DB) FeedPushEventDAO {
	return &GORMFeedPushEventDAO{db: db}
}

func (dao *GORMFeedPushEventDAO) FindFeedEvents(ctx context.Context, uid int64, lastTime int64, direction int32,
	limit int64) ([]FeedPushEvent, error) {
	query := dao.db.WithContext(ctx).Where("uid = ?", uid, lastTime)
	const (
		DirectionBefore = 0
		DirectionAfter  = 1
	)
	switch direction {
	case DirectionBefore:
		query = query.Where("ctime < ?").Order("ctime DESC")
	case DirectionAfter:
		query = query.Where("ctime > ?").Order("ctime ASC")
	default:
		return nil, errors.New("不合法的查询方向")
	}
	var events []FeedPushEvent
	err := query.Limit(int(limit)).Find(&events).Error
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
