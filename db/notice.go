package db

import (
	"boot/ent"
	"boot/ent/notice"
	"context"
)

// FindNoticeByUserIdAndType 通过用户id和类型获取信息
func FindNoticeByUserIdAndType(ctx context.Context, userId int, noticeType int) ([]*ent.Notice, error) {
	all, err := client.Notice.Query().Where(notice.UserId(userId), notice.NoticeType(noticeType)).All(ctx)
	return all, err
}
