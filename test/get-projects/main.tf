terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
    }
  }
}

provider "platformsh" {
  api_token = "NLiGXU8Z6HzIkMjXHl6PKtzlXIflIF4xgWio63OjSvc"
}

data "platformsh_projects" "example" {}

output "project_ids" {
  value = [for p in data.platformsh_projects.example.projects : p.id]
}
