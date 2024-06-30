// main.tf
terraform {
  required_providers {
    hashicups = {
      source = "local.provider/rhs/platformsh"
    }
  }
}

provider "hashicups" {}

data "hashicups_coffees" "example" {}
