terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
      version = "0.1.0"
    }
  }
}

provider "platformsh" {
  api_token = "YOUR_API_KEY"
}

resource "platformsh_environment" "new_environment" {
  project_id      = "PROJECT_ID"
  name            = "test-env"
  title           = "test-environment"
  }

output "environment_status" {
  value = platformsh_environment.new_environment.status
}
