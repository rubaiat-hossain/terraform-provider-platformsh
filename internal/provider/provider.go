package provider

import (
	"context"

	"github.com/rubaiat-hossain/terraform-provider-platformsh/internal/platformsh"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &platformshProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &platformshProvider{
			version: version,
		}
	}
}

// platformshProvider defines the provider implementation
type platformshProvider struct {
	version string
}

// Metadata returns the provider type name
func (p *platformshProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "platformsh"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data
func (p *platformshProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				Description: "API token for Platform.sh",
				Required:    true,
			},
		},
	}
}

// Configure prepares the API client
func (p *platformshProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config struct {
		APIToken types.String `tfsdk:"api_token"`
	}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new Platform.sh client using the configuration values
	client, err := platformsh.NewClient(config.APIToken.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Platform.sh API Client",
			"An unexpected error occurred when creating the Platform.sh API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Platform.sh Client Error: "+err.Error(),
		)
		return
	}

	// Make the Platform.sh client available during DataSource and Resource type Configure methods
	resp.DataSourceData = client
	resp.ResourceData = client
}

// Resources returns the resource implementations
func (p *platformshProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPlatformshEnvironmentResource,
	}
}
