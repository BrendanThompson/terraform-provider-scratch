package scratch

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &BlockResource{}

func NewBlockResource() resource.Resource {
	return &BlockResource{}
}

type BlockResource struct{}

type BlockResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	In          []BlockProps `tfsdk:"in"`
}

type BlockProps struct {
	String types.String `tfsdk:"string"`
	Number types.Number `tfsdk:"number"`
	Bool   types.Bool   `tfsdk:"bool"`
}

func (r *BlockResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_block"
}

func (r *BlockResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
		},
		Blocks: map[string]tfsdk.Block{
			"in": {
				Attributes: map[string]tfsdk.Attribute{
					"string": {
						Type:     types.StringType,
						Optional: true,
					},
					"number": {
						Type:     types.NumberType,
						Optional: true,
					},
					"bool": {
						Type:     types.BoolType,
						Optional: true,
					},
				},
				MinItems:    0,
				NestingMode: 2,
			},
		},
	}, nil
}

func (r *BlockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BlockResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	id, err := uuid.NewUUID()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to generate UUID",
			err.Error(),
		)
		return
	}

	data.ID = types.String{Value: id.String()}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BlockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BlockResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BlockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan BlockResourceModel
	var state BlockResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	var result = BlockResourceModel{
		ID:          state.ID,
		Description: plan.Description,
		In:          plan.In,
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)
}

func (r *BlockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BlockResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	resp.State.RemoveResource(ctx)
}
