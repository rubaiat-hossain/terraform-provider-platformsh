terraform {
  required_providers {
    platformsh = {
      source = "local.provider/rhs/platformsh"
      version = "0.1.0"
    }
  }
}

provider "platformsh" {
  api_token = "<YOUR-API-KEY>"
}

resource "platformsh_environment" "new_environment" {
  project_id      = "<PROJECT-ID>"
  name            = "latest-env"
  title           = "Latest Environment"
  }

output "environment_status" {
  value = platformsh_environment.new_environment.status
}
