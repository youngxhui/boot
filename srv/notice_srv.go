package srv

import (
	"boot/protos"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NoticeService struct {
}

func (n NoticeService) ListNotices(ctx context.Context, in *protos.ListNoticesRequest) (*protos.ListNoticesResponse, error) {
	return nil, nil
}
func (n NoticeService) GetNotice(ctx context.Context, in *protos.GetNoticeRequest) (*protos.Notice, error) {
	return nil, nil
}
func (n NoticeService) CreateNotice(ctx context.Context, in *protos.CreateNoticeRequest) (*protos.Notice, error) {
	return nil, nil
}
func (n NoticeService) UpdateNotice(ctx context.Context, in *protos.UpdateNoticeRequest) (*protos.Notice, error) {
	return nil, nil
}
func (n NoticeService) DeleteNotice(ctx context.Context, in *protos.DeleteNoticeRequest) (*emptypb.Empty, error) {
	return nil, nil
}
