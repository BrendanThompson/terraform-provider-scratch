package scratch

import "github.com/hashicorp/terraform-plugin-framework/types"

type String struct {
	ID          types.String `tfsdk:"id"`
	In          types.String `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}

type List struct {
	ID          types.String `tfsdk:"id"`
	In          types.List   `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}

type Number struct {
	ID          types.String `tfsdk:"id"`
	In          types.Number `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}

type Bool struct {
	ID          types.String `tfsdk:"id"`
	In          types.Bool   `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}

type Map struct {
	ID          types.String `tfsdk:"id"`
	In          types.Map    `tfsdk:"in"`
	Description types.String `tfsdk:"description"`
}
