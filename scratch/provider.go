package scratch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &ScratchProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ScratchProvider{
			Version: version,
		}
	}
}

type ScratchProvider struct {
	Version string

	isConfigured bool
}

type ScratchProviderModel struct {
	Use types.Bool `tfsdk:"use"`
}

func (p *ScratchProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scratch"
}

func (p *ScratchProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"use": {
				Type:     types.BoolType,
				Optional: true,
			},
		},
	}, nil
}

func (p *ScratchProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data ScratchProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if !data.Use.Value {
		data.Use = types.Bool{Value: true}
	}

	p.isConfigured = true
}

func (p *ScratchProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewStringDataSource,
	}
}

func (p *ScratchProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewBlockResource,
		NewBoolResource,
		NewListResource,
		NewMapResource,
		NewNumberResource,
		NewStringResource,
	}
}
