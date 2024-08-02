terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
      version = "0.1.0"
    }
  }
}

provider "platformsh" {
  api_token = "NLiGXU8Z6HzIkMjXHl6PKtzlXIflIF4xgWio63OjSvc"
}

resource "platformsh_environment" "new_environment" {
  project_id      = "roqsqouvgnwsm"
  name            = "lates-env"
  title           = "Latest Environment"
  }

output "environment_status" {
  value = platformsh_environment.new_environment.status
}
