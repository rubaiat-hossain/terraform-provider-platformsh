package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rubaiat-hossain/terraform-provider-platformsh/internal/platformsh"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &platformshEnvironmentResource{}
var _ resource.ResourceWithConfigure = &platformshEnvironmentResource{}
var _ resource.ResourceWithImportState = &platformshEnvironmentResource{}

func NewEnvironmentResource() resource.Resource {
	return &platformshEnvironmentResource{}
}

// platformshEnvironmentResource defines the resource implementation.
type platformshEnvironmentResource struct {
	client *platformsh.Client
}

// platformshEnvironmentResourceModel describes the resource data model.
type platformshEnvironmentResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	ProjectID             types.String `tfsdk:"project_id"`
	EnvironmentName       types.String `tfsdk:"environment_name"`
	ConfigurableAttribute types.String `tfsdk:"configurable_attribute"`
	Defaulted             types.String `tfsdk:"defaulted"`
}

func (r *platformshEnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment"
}

func (r *platformshEnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Platform.sh environment resource",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Description: "ID of the Platform.sh project.",
				Required:    true,
			},
			"environment_name": schema.StringAttribute{
				Description: "Name of the environment.",
				Required:    true,
			},
			"configurable_attribute": schema.StringAttribute{
				Description: "Example configurable attribute",
				Optional:    true,
			},
			"defaulted": schema.StringAttribute{
				Description: "Example configurable attribute with default value",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("example value when not configured"),
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *platformshEnvironmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*platformsh.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *platformsh.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *platformshEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data platformshEnvironmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to create the environment
	// For the purposes of this example, hardcoding a response value to save into the Terraform state.
	data.ID = types.StringValue("example-id")

	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *platformshEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data platformshEnvironmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to read the environment
	// For the purposes of this example, assume the environment is still available and up-to-date.

	tflog.Trace(ctx, "read a resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *platformshEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data platformshEnvironmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to update the environment
	// For the purposes of this example, hardcoding a response value to save into the Terraform state.
	data.ID = types.StringValue("example-id")

	tflog.Trace(ctx, "updated a resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *platformshEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data platformshEnvironmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to delete the environment
	// For the purposes of this example, assume the environment was successfully deleted.

	tflog.Trace(ctx, "deleted a resource")

	// Remove resource from Terraform state
	resp.State.RemoveResource(ctx)
}

func (r *platformshEnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve the resource ID from the import request
	resourceID := req.ID

	// Set the resource ID in the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &platformshEnvironmentResourceModel{
		ID: types.StringValue(resourceID),
	})...)
}
