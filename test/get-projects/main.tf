terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
    }
  }
}

provider "platformsh" {
  api_token = "<YOUR-API-KEY>"
}

data "platformsh_projects" "example" {}

output "project_ids" {
  value = [for p in data.platformsh_projects.example.projects : p.id]
}
