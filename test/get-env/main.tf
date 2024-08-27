terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
    }
  }
}

provider "platformsh" {
  api_token = "YOUR_API_KEY"
}

data "platformsh_environments" "example" {
  project_id = "PROJECT_ID"
}

output "environment_ids" {
  value = [for e in data.platformsh_environments.example.environments : e.id]
}

