package scratch

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &StringDataSource{}

func NewStringDataSource() datasource.DataSource {
	return &StringDataSource{}
}

type StringDataSource struct{}

type StringDataSourceModel struct {
	ID  types.String `tfsdk:"id"`
	In  types.String `tfsdk:"in"`
	Out types.String `tfsdk:"out"`
}

func (s *StringDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "scratch_string"
}

func (s StringDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"in": {
				Type:     types.StringType,
				Required: true,
			},
			"out": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

func (s *StringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data StringDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	uuid, err := uuid.NewUUID()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to generate UUID",
			err.Error(),
		)
		return
	}

	var result = StringDataSourceModel{
		ID:  types.String{Value: uuid.String()},
		In:  data.In,
		Out: data.In,
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)

}
