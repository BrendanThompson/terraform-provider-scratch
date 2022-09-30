package scratch

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &NumberResource{}

func NewNumberResource() resource.Resource {
	return &NumberResource{}
}

type NumberResource struct{}

type NumberResourceModel struct {
	ID          types.String `tfsdk:"id"`
	In          types.Number `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}

func (r *NumberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_number"
}

func (r *NumberResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"in": {
				Type:     types.NumberType,
				Required: true,
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (r *NumberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NumberResourceModel

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

func (r *NumberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NumberResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NumberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NumberResourceModel
	var state NumberResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	var result = NumberResourceModel{
		ID:          state.ID,
		Description: plan.Description,
		In:          plan.In,
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)
}

func (r *NumberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NumberResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	resp.State.RemoveResource(ctx)
}
