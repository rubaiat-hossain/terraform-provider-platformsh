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

data "platformsh_environments" "example" {
  project_id = "roqsqouvgnwsm"
}

output "environment_ids" {
  value = [for e in data.platformsh_environments.example.environments : e.id]
}

