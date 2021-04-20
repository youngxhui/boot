package srv

import (
	"boot/entity"
	pd "boot/protos"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ToolService struct {
}

// ListTools 获取 10 个数据
func (t *ToolService) ListTools(ctx context.Context, in *pd.ListToolsRequest) (*pd.ListToolsResponse, error) {
	size := in.PageSize
	tools := entity.FindAllTools(int(size))
	pdTools := make([]*pd.Tool, len(tools))
	for i := 0; i < len(tools); i++ {
		pdTools[i] = &pd.Tool{
			Id:        int32(tools[i].ID),
			MachineId: int32(tools[i].MachineId),
			Status:    pd.Status(tools[i].Status),
		}
	}
	return &pd.ListToolsResponse{Tools: pdTools}, nil
}

// GetTool 通过 id 获取刀具状态
func (t *ToolService) GetTool(ctx context.Context, in *pd.GetToolRequest) (*pd.Tool, error) {
	tool := entity.Tool{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
	}
	toolFirst := tool.FindById()
	if toolFirst == nil {
		return nil, status.Error(codes.ResourceExhausted, "没有该刀具信息")
	}
	return &pd.Tool{
		Id:        int32(toolFirst.ID),
		MachineId: int32(toolFirst.MachineId),
		Status:    pd.Status(toolFirst.Status),
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
