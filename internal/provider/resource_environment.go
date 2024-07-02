package provider

import (
	"context"

	"github.com/rubaiat-hossain/terraform-provider-platformsh/internal/platformsh"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ resource.Resource = &platformshEnvironmentResource{}
)

type platformshEnvironmentResource struct {
	client *platformsh.Client
}

func NewPlatformshEnvironmentResource() resource.Resource {
	return &platformshEnvironmentResource{}
}

// Schema defines the schema for the resource
func (r *platformshEnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource.Schema{
		Attributes: map[string]resource.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"status": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}
}

// Configure prepares the resource with the configured client
func (r *platformshEnvironmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if v, ok := req.ProviderData.(*platformsh.Client); ok {
		r.client = v
	}
}

// Create a new environment
func (r *platformshEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Implementation here
}

// Read an existing environment
func (r *platformshEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Implementation here
}

// Update an existing environment
func (r *platformshEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Implementation here
}

// Delete an environment
func (r *platformshEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Implementation here
}
