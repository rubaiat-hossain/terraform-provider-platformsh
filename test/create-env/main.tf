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

resource "platformsh_environment" "example" {
  project_id     = "roqsqouvgnwsm"
  name           = "example-env"
  title          = "Example Environment"
  enable_smtp    = true
  restrict_robots = false
}

output "environment_id" {
  value = platformsh_environment.example.id
}

