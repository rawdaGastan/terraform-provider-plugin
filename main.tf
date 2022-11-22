terraform {
  required_providers {
    plugin = {
      source  = "example.com/local/plugin"
      version = "~> 1.0.0"
    }
  }
}

provider "plugin" {
  url = "http://localhost:3000"
}

resource "provider_query" "vm_1" {
  provider = plugin

  project = "pkid"
  key = "key"
  value = "value"
  encrypt = true
}

output "vm_1" {
  value = provider_query.vm_1
}