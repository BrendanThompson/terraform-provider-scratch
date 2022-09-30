package scratch

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &StringResource{}

func NewStringResource() resource.Resource {
	return &StringResource{}
}

type StringResource struct{}

type StringResourceModel struct {
	ID          types.String `tfsdk:"id"`
	In          types.String `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}

func (r *StringResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_string"
}

func (r *StringResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (r *StringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data StringResourceModel

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

func (r *StringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data StringResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *StringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan StringResourceModel
	var state StringResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	var result = StringResourceModel{
		ID:          state.ID,
		Description: plan.Description,
		In:          plan.In,
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)
}

func (r *StringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data StringResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	resp.State.RemoveResource(ctx)
}
