terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
    }
  }
}

provider "platformsh" {
  api_token = "bs0oAldrnjRH6JMPSLtqrwjKzlgLudvXXc7Es8Zo2lQ"
}

data "platformsh_projects" "example" {}

output "project_ids" {
  value = [for p in data.platformsh_projects.example.projects : p.id]
}