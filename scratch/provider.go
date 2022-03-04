package scratch

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var stderr = os.Stderr

type provider struct {
	configured bool
	client     string //TODO: use client
}

type providerData struct {
	Use types.Bool `tfsdk:"use"`
}

func New() tfsdk.Provider {
	return &provider{}
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !config.Use.Value {
		config.Use = types.Bool{Value: true}
	}

	p.configured = true
}

func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"use": {
				Type:     types.BoolType,
				Optional: true,
			},
		},
	}, nil
}

func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"scratch_string": resourceStringType{},
		"scratch_number": resourceNumberType{},
		"scratch_bool":   resourceBoolType{},
		"scratch_list":   resourceListType{},
		"scratch_map":    resourceMapType{},
		"scratch_block":  resourceBlockType{},
	}, nil
}

func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{}, nil
}
