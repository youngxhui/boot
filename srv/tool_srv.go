package srv

import (
	"boot/entity"
	pd "boot/protos"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ToolService struct {
}

func (t *ToolService) ListTools(ctx context.Context, in *pd.ListToolsRequest) (*pd.ListToolsResponse, error) {
	return nil, nil
}
func (t *ToolService) GetTool(ctx context.Context, in *pd.GetToolRequest) (*pd.Tool, error) {
	tool := entity.Tool{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
	}
	toolFirst := tool.FindById()

	return &pd.Tool{
		Id:        toolFirst.Id,
		MachineId: toolFirst.MachineId,
		Status:    toolFirst.Status,
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
