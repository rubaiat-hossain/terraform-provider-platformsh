package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/rubaiat-hossain/terraform-provider-platformsh/internal/platformsh"
)

// Ensure provider implementation satisfies the provider.Provider interface.
var _ provider.Provider = &platformshProvider{}

// New creates a new platformsh provider.
func New() provider.Provider {
	return &platformshProvider{}
}

// platformshProvider defines the provider implementation.
type platformshProvider struct {
	client *platformsh.Client
}

// Metadata returns the provider type name.
func (p *platformshProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "platformsh"
}

// Schema defines the provider-level schema for configuration data.
func (p *platformshProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Platform.sh provider",
	}
}

// Configure prepares the provider with the given configuration.
func (p *platformshProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Configuration data can be retrieved from req.Config
	// For simplicity, we assume a static base URL and use http.DefaultClient
	baseURL := "https://api.platform.sh"
	p.client = platformsh.NewClient(baseURL, http.DefaultClient)
}

// Resources returns the resource implementations supported by this provider.
func (p *platformshProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewEnvironmentResource,
	}
}

// DataSources returns the data source implementations supported by this provider.
func (p *platformshProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewEnvironmentDataSource,
	}
}
