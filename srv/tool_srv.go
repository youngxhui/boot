package srv

import (
	"boot/db"
	pd "boot/protos"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ToolService struct {
}

// ListTools 获取 10 个数据
func (t *ToolService) ListTools(ctx context.Context, in *pd.ListToolsRequest) (*pd.ListToolsResponse, error) {

	allTools, err := db.FindAllTools(ctx, int(in.Page), int(in.Size))
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	pdTools := make([]*pd.Tool, len(allTools))
	for i := 0; i < len(allTools); i++ {
		pdTools[i] = &pd.Tool{
			Id:        int32(allTools[i].ID),
			MachineId: int32(allTools[i].MachineID),
			Status:    pd.Status(allTools[i].Status),
		}
	}
	return &pd.ListToolsResponse{Tools: pdTools}, nil
}

// GetTool 通过 id 获取刀具状态
func (t *ToolService) GetTool(ctx context.Context, in *pd.GetToolRequest) (*pd.Tool, error) {
	tool, err := db.FindToolById(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.ResourceExhausted, err.Error())
	}

	return &pd.Tool{
		Id:        int32(tool.ID),
		MachineId: int32(tool.MachineID),
		Status:    pd.Status(tool.Status),
	}, nil
}

func (t *ToolService) CreateTool(ctx context.Context, in *pd.CreateToolRequest) (*pd.Tool, error) {

	return nil, nil
}
func (t *ToolService) UpdateTool(ctx context.Context, in *pd.UpdateToolRequest) (*pd.Tool, error) {
	return nil, nil
}
func (t *ToolService) DeleteTool(ctx context.Context, in *pd.DeleteToolRequest) (*emptypb.Empty, error) {
	return nil, nil
}
