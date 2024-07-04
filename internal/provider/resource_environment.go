package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rubaiat-hossain/terraform-provider-platformsh/internal/platformsh"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EnvironmentResource{}
var _ resource.ResourceWithConfigure = &EnvironmentResource{}
var _ resource.ResourceWithImportState = &EnvironmentResource{}

func NewEnvironmentResource() resource.Resource {
	return &EnvironmentResource{}
}

// EnvironmentResource defines the resource implementation.
type EnvironmentResource struct {
	client *platformsh.Client
}

// EnvironmentResourceModel describes the resource data model.
type EnvironmentResourceModel struct {
	ID             types.String `tfsdk:"id"`
	ProjectID      types.String `tfsdk:"project_id"`
	Name           types.String `tfsdk:"name"`
	Title          types.String `tfsdk:"title"`
	Type           types.String `tfsdk:"type"`
	Status         types.String `tfsdk:"status"`
	DefaultDomain  types.String `tfsdk:"default_domain"`
	EnableSMTP     types.Bool   `tfsdk:"enable_smtp"`
	RestrictRobots types.Bool   `tfsdk:"restrict_robots"`
}

func (r *EnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment"
}

func (r *EnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the environment",
				Computed:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "ID of the project",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the environment",
				Required:    true,
			},
			"title": schema.StringAttribute{
				Description: "Title of the environment",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "Type of the environment",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the environment",
				Computed:    true,
			},
			"default_domain": schema.StringAttribute{
				Description: "Default domain of the environment",
				Computed:    true,
			},
			"enable_smtp": schema.BoolAttribute{
				Description: "Enable SMTP for the environment",
				Optional:    true,
				Computed:    true,
			},
			"restrict_robots": schema.BoolAttribute{
				Description: "Restrict robots for the environment",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (r *EnvironmentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *EnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EnvironmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to create the environment
	environment := &platformsh.Environment{
		Name:           data.Name.ValueString(),
		Title:          data.Title.ValueString(),
		EnableSMTP:     data.EnableSMTP.ValueBool(),
		RestrictRobots: data.RestrictRobots.ValueBool(),
	}

	createdEnvironment, err := r.client.CreateEnvironment(data.ProjectID.ValueString(), environment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			"Unable to create environment, got error: "+err.Error(),
		)
		return
	}

	// Save data into Terraform state
	data.ID = types.StringValue(createdEnvironment.ID)
	data.Type = types.StringValue(createdEnvironment.Type)
	data.Status = types.StringValue(createdEnvironment.Status)
	data.DefaultDomain = types.StringValue(createdEnvironment.DefaultDomain)
	data.EnableSMTP = types.BoolValue(createdEnvironment.EnableSMTP)
	data.RestrictRobots = types.BoolValue(createdEnvironment.RestrictRobots)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the Read function to sync the state
	readReq := resource.ReadRequest{
		State: resp.State,
	}
	readResp := &resource.ReadResponse{}
	r.Read(ctx, readReq, readResp)
	resp.Diagnostics.Append(readResp.Diagnostics...)
}

func (r *EnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data EnvironmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to read the environment
	environment, err := r.client.GetEnvironment(data.ProjectID.ValueString(), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			"Unable to read environment, got error: "+err.Error(),
		)
		return
	}

	// Save updated data into Terraform state
	data.Name = types.StringValue(environment.Name)
	data.Title = types.StringValue(environment.Title)
	data.Type = types.StringValue(environment.Type)
	data.Status = types.StringValue(environment.Status)
	data.DefaultDomain = types.StringValue(environment.DefaultDomain)
	data.EnableSMTP = types.BoolValue(environment.EnableSMTP)
	data.RestrictRobots = types.BoolValue(environment.RestrictRobots)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data EnvironmentResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to update the environment
	environment := &platformsh.Environment{
		Name:           data.Name.ValueString(),
		Title:          data.Title.ValueString(),
		EnableSMTP:     data.EnableSMTP.ValueBool(),
		RestrictRobots: data.RestrictRobots.ValueBool(),
	}

	updatedEnvironment, err := r.client.UpdateEnvironment(data.ProjectID.ValueString(), data.ID.ValueString(), environment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			"Unable to update environment, got error: "+err.Error(),
		)
		return
	}

	// Save updated data into Terraform state
	data.Type = types.StringValue(updatedEnvironment.Type)
	data.Status = types.StringValue(updatedEnvironment.Status)
	data.DefaultDomain = types.StringValue(updatedEnvironment.DefaultDomain)
	data.EnableSMTP = types.BoolValue(updatedEnvironment.EnableSMTP)
	data.RestrictRobots = types.BoolValue(updatedEnvironment.RestrictRobots)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call the Read function to sync the state
	readReq := resource.ReadRequest{
		State: resp.State,
	}
	readResp := &resource.ReadResponse{}
	r.Read(ctx, readReq, readResp)
	resp.Diagnostics.Append(readResp.Diagnostics...)
}

func (r *EnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data EnvironmentResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Call Platform.sh API to delete the environment
	err := r.client.DeleteEnvironment(data.ProjectID.ValueString(), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			"Unable to delete environment, got error: "+err.Error(),
		)
		return
	}

	// Remove resource from Terraform state
	resp.State.RemoveResource(ctx)
}

func (r *EnvironmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve the resource ID from the import request
	resourceID := req.ID

	// Set the resource ID in the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &EnvironmentResourceModel{
		ID: types.StringValue(resourceID),
	})...)
}
