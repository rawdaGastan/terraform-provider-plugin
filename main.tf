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
  seed = "bm2xl92552zz0Kxtvg4Gbaosnh6FY9H2WsIKao6Emh8="
}

resource "plugin_pkid_key_query" "vm_1" {
  provider = plugin

  project = "pkid"
  key = "key"
  value = "value"
  encrypt = true
}

output "vm_1" {
  value = plugin_pkid_key_query.vm_1
}

resource "plugin_pkid_project_query" "vm_2" {
  provider = plugin

  project = "pkid"
  key = "key"
  value = "value"
  encrypt = true
}

output "vm_2" {
  value = plugin_pkid_project_query.vm_2
}
